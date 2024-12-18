package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	consumer "ms-metawebhooks/consumer"
	webhook_data_entity "ms-metawebhooks/webhook_data_entity"
	whatsapp "ms-metawebhooks/whatsapp"

	"github.com/joho/godotenv"
)

func HandleWhatsappPhoneRoute(verifier whatsapp.Verifier, receiver whatsapp.Receiver) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			mode := r.URL.Query().Get("hub.mode")
			challenge := r.URL.Query().Get("hub.challenge")
			verify_token := r.URL.Query().Get("hub.verify_token")

			is_valid := verifier.Verify(mode, verify_token)

			if is_valid {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(challenge))
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request!"))
		case http.MethodPost:
			phone := r.URL.Path[len("/api/v1/webhooks/whatsapp/"):]
			body, err := io.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Bad Request!"))
			}

			receiver.Receive(phone, body)
      webhook_data_entity.Create(phone, string(body))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed!"))
		}
	}
}

func main() {
	err := godotenv.Load("../.env")
	
	if err != nil {
		fmt.Println("Error loading .env file")
	}

  go consumer.Handle()

	port := os.Getenv("PORT")
	fmt.Println("Server running on port: ", port)

	http.HandleFunc("/api/v1/webhooks/whatsapp/{phone}", HandleWhatsappPhoneRoute(whatsapp.DefaultVerifier{}, whatsapp.DefaultReceiver{}))

	http.ListenAndServe(":"+port, nil)
}
