package publisher

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Handle(phone string, body []byte) {
	rabbitmq_uri := os.Getenv("RABBITMQ_URI")

	conn, err := amqp.Dial(rabbitmq_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"whatsapp_webhooks", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	api_access_token := os.Getenv("API_ACCESS_TOKEN")

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        body,
			Headers: amqp.Table {
				"api_access_token": api_access_token,
				"phone": phone,
			},
		},
	)

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}