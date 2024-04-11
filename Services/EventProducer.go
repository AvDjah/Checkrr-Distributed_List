package Services

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

type Message struct {
	EventId   int64
	EventType int
	Message   string
	UserId    int64
}

var DbChannel chan Message

var NotificationChannel chan Message

var Conn *amqp091.Connection

func initDbConnection() chan Message {
	if DbChannel != nil {
		return DbChannel
	} else {
		DbChannel = make(chan Message, 1e5)
		return DbChannel
	}
}

func initNotificationChannel() chan Message {
	if NotificationChannel != nil {
		return NotificationChannel
	} else {
		NotificationChannel = make(chan Message, 1e5)
		return NotificationChannel
	}
}

func GetConnection() *amqp091.Connection {
	if Conn != nil {
		return Conn
	} else {
		conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
		failOnError(err, "Dialing AMQP")
		return conn
	}
}

func RunWorker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Run The Two Consumers
	go testRunDbConsumer()
	go testRunNotificationConsumer()

	for {
		select {
		case <-ticker.C:
			if time.Now().Second() == 55 {
				go worker()
			}
		}
	}
}

func worker() {
	db := Db.InitDb()
	events, _ := Data.GetEventsForNextMinute(db)
	var triggeredEvents []Models.Event
	for _, event := range events {
		// Add a check to get the event which we are going to happen
		triggeredEvents = append(triggeredEvents, event)
	}

	// Declare Channels
	dbChannel := initDbConnection()
	notificationChannel := initNotificationChannel()

	for _, item := range triggeredEvents {

		// Send Triggered Notifications to be notified to Users in Real Time
		dbChannel <- Message{
			EventType: item.EventType,
			Message:   item.Description,
			EventId:   item.ID,
			UserId:    item.UserID,
		}

		// Send Triggered Notifications to be stored in DB
		notificationChannel <- Message{
			EventType: item.EventType,
			Message:   item.Description,
			EventId:   item.ID,
			UserId:    item.UserID,
		}
		produceEvent(Message{
			EventType: item.EventType,
			Message:   item.Description,
			EventId:   item.ID,
			UserId:    item.UserID,
		})
	}

}

func produceEvent(msg Message) {
	conn := GetConnection()
	defer func(conn *amqp091.Connection) {
		err := conn.Close()
		if err != nil {
			failOnError(err, "Error Closing")
		}
	}(conn)

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func(ch *amqp091.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Closing Channel")
		}
	}(ch)

	err = ch.ExchangeDeclare(
		"checkrr", // name
		"fanout",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Declaring Exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	jsonData, err := json.Marshal(gin.H{
		"data": msg,
	})
	failOnError(err, "Marshalling JSON")

	err = ch.PublishWithContext(ctx,
		"checkrr", // exchange
		"",        // routing key
		false,     // mandatory
		false,     // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(jsonData),
		})
}
