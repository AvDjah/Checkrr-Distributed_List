package Services

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func RunDbConsumer() {
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
			fmt.Println("DB Receiver received: ", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func testRunDbConsumer() {
	dbChannel := initDbConnection()

	for {
		select {
		case msg := <-dbChannel:
			fmt.Println("Received Event: ID: ", msg.EventId, "Type : ", msg.EventType)
		}
	}

}
