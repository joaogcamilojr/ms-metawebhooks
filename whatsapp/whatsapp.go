package whatsapp

import (
	"bytes"
	"fmt"
	publisher "ms-webhooks/publisher"
	"net/http"
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

func Send(body []byte) {
    token := ""
    cloud_api_base_url := ""

    req, err := http.NewRequest("POST", cloud_api_base_url, bytes.NewBuffer(body))

    if err != nil {
        fmt.Println("Error on create request", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        fmt.Println("Error on send request", err)
        return
    }

    defer resp.Body.Close()

    fmt.Println("Status Code: ", resp.StatusCode)

    // Read Response
}