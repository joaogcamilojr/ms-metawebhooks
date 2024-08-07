package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockVerifier struct {}

type MockReceiver struct {}

func (v MockVerifier) Verify(mode string, verify_token string) (bool) {
	return mode == "subscribe" && verify_token == "my_verify_token"
}

func (r MockReceiver) Receive(phone string, body []byte) {
	if phone != "5535987654321"|| string(body) != "test body content" {
		panic("unexpected phone or body content")
	}
}

func TestHandleWhatsappPhoneRouteGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/webhooks/whatsapp/5535987654321?hub.mode=subscribe&hub.challenge=1234&hub.verify_token=my_verify_token", nil)
	rec := httptest.NewRecorder()
	
	handler := HandleWhatsappPhoneRoute(MockVerifier{}, MockReceiver{})

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", res.StatusCode)
	}

	body, _ := io.ReadAll(res.Body)

	if string(body) != "1234" {
		t.Errorf("expected body '1234' got %s", string(body))
	}
}

func TestHandleWhatsappPhoneRoutePost(t *testing.T) {
	body := strings.NewReader("test body content")
	req := httptest.NewRequest(http.MethodPost, "/api/v1/webhooks/whatsapp/5535987654321", body)
	rec := httptest.NewRecorder()
	
	handler := HandleWhatsappPhoneRoute(MockVerifier{}, MockReceiver{})

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", res.StatusCode)
	}
}

func TestHandleWhatsappPhoneRouteInvalidRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/webhooks/whatsapp/5535987654321", nil)
	rec := httptest.NewRecorder()
	
	handler := HandleWhatsappPhoneRoute(MockVerifier{}, MockReceiver{})

	handler.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %v", res.StatusCode)
	}
}