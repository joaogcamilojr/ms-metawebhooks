package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	whatsapp "ms-webhooks/whatsapp"
)

func HandleWhatsappPhoneRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mode := r.URL.Query().Get("hub.mode")
		challenge := r.URL.Query().Get("hub.challenge")
		verify_token := r.URL.Query().Get("hub.verify_token")

		is_valid := whatsapp.Verify(mode, verify_token)

		if is_valid {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(challenge))
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request!"))
	case http.MethodPost:
		phone := r.PathValue("phone")
		body, err := io.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request!"))
		}

		whatsapp.Receive(phone, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed!"))
	}
}

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Server running on port: ", port)

	http.HandleFunc("/api/v1/webhooks/whatsapp/{phone}", HandleWhatsappPhoneRoute)

	http.ListenAndServe(":"+port, nil)
}