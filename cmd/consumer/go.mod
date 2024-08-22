module ms-webhooks/consumer

go 1.22.5

require github.com/rabbitmq/amqp091-go v1.10.0

replace ms-webhooks/whatsapp => ./whatsapp

replace ms-webhooks/publisher => ./publisher
