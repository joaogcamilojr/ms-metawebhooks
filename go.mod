module joaogcamilojr/ms-webhooks

go 1.22.5

replace ms-webhooks/whatsapp => ./whatsapp

replace ms-webhooks/publisher => ./publisher

require (
	github.com/joho/godotenv v1.5.1
	ms-webhooks/consumer v0.0.0-00010101000000-000000000000
	ms-webhooks/whatsapp v0.0.0-00010101000000-000000000000
)

require (
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	ms-webhooks/publisher v0.0.0-00010101000000-000000000000 // indirect
)

replace ms-webhooks/consumer => ./consumer
