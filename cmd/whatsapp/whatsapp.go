package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	publisher "ms-metawebhooks/publisher"
	"net/http"
	"os"
)

type NewStructure struct {
    Pattern string      `json:"pattern"`
    Data    interface{} `json:"data"`
}

type DefaultVerifier struct {}

type DefaultReceiver struct {}

type Verifier interface {
	Verify(mode string, verify_token string) (bool)
}

type Receiver interface {
	Receive(phone string, body []byte)
}


func (v DefaultVerifier) Verify(mode string, verify_token string) (bool) {
	valid_verify_token := os.Getenv("VERIFY_TOKEN")

	return mode == "subscribe" && verify_token == valid_verify_token
}

func (r DefaultReceiver) Receive(phone string, body []byte) {
	fmt.Println("phone: ", phone)

    var originalBody map[string]interface{}

    err := json.Unmarshal(body, &originalBody)

    if (err != nil) {
        fmt.Errorf("Error parsing body: %w", err)
    }

    newBody := NewStructure {
        Pattern: "whatsapp",
        Data: originalBody,
    }

    result, err := json.Marshal(newBody)


    if (err != nil) {
        fmt.Errorf("Error creating new body: %w", err)
    }

	publisher.Handle(phone, result)
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
