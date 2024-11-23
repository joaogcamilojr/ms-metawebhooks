module joaogcamilojr/ms-metawebhooks

go 1.22.5

replace ms-metawebhooks/whatsapp => ./whatsapp

replace ms-metawebhooks/publisher => ./publisher

require (
	github.com/joho/godotenv v1.5.1
	ms-metawebhooks/consumer v0.0.0-00010101000000-000000000000
	ms-metawebhooks/whatsapp v0.0.0-00010101000000-000000000000
)

require (
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	ms-metawebhooks/publisher v0.0.0-00010101000000-000000000000 // indirect
)

replace ms-metawebhooks/consumer => ./consumer
