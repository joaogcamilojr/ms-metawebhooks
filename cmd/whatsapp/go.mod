module ms-webhooks/whatsapp

go 1.22.2

replace ms-webhooks/publisher => ../publisher

require ms-webhooks/publisher v0.0.0-00010101000000-000000000000

require github.com/rabbitmq/amqp091-go v1.10.0 // indirect
