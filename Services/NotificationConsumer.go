package Services

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func HandleNotification(msg amqp091.Delivery) {
	// Handle Delivery to Client
}

func RunNotificationConsumer() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Dialing AMQP Server")

	ch, err := conn.Channel()
	failOnError(err, "Creating AMQP Channel")
	defer func(ch *amqp091.Channel) {
		err := ch.Close()
		if err != nil {
			failOnError(err, "Closing Channel")
		}
	}(ch)

	q, err := ch.QueueDeclare("checkrr", false, false, false, false, nil)

	failOnError(err, "Declaring Queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			HandleNotification(d)
			fmt.Println("Notification Receiver received: ", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func testRunNotificationConsumer() {
	notificationChannel := initNotificationChannel()

	for {
		select {
		case msg := <-notificationChannel:
			fmt.Println("Received Event: ID: ", msg.EventId, "Type : ", msg.EventType)
		}
	}
}
