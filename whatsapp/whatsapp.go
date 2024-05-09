package whatsapp

import (
	"fmt"
	publisher "ms-webhooks/publisher"
	"os"
)

func Verify(mode string, verify_token string) (bool) {
	valid_verify_token := os.Getenv("VERIFY_TOKEN")

	return mode == "subscribe" && verify_token == valid_verify_token
}

func Receive(phone string, body []byte) {
	fmt.Println("phone: ", phone)

	publisher.Handle(phone, body)
}