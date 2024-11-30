package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	whatsapp "ms-metawebhooks/whatsapp"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageBody struct {
    Product string `json:"pattern"`
    Message map[string]interface{} `json:"data"`
}

func failOnError(err error, msg string) {
    if err != nil {
    log.Panicf("%s: %s", msg, err)
    }
}


func Handle() {
	  rabbitmq_uri := os.Getenv("RABBITMQ_URI")

    conn, err := amqp.Dial(rabbitmq_uri)
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "meta_outputs", // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

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
            log.Printf("Received a message: %s", d.Body)

            var message MessageBody

            err := json.Unmarshal(d.Body, &message)

            if err != nil {
                log.Fatal("Error unmarshalling json", err)
            }

            body, err := json.Marshal(message.Message)

            if err != nil {
                log.Fatal("Error taking body", err)
            }

            switch message.Product {
                case "whatsapp": 
                    fmt.Println("Sendind whatsapp message")
                    whatsapp.Send(body)
                default:
                    fmt.Println("Product not implemented: ", message.Product)
            }
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
}
