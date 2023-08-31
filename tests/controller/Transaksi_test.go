package tests

import (
	"link-aja/controller/transaksi"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func AddBalance(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/account/{number_account}/top-up", transaksi.AddBalance)

	payload := strings.NewReader(`{"balance": 10000}`)

	req, err := http.NewRequest("POST", "http://localhost:5000/account/0000001/top-up", payload)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected status %v, but got %v", http.StatusCreated, response.Code)
	}
}

func Trasnfer(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/account/{number_account}/transfer", transaksi.Transfer)

	payload := strings.NewReader(`{"to_account_number": "0000002", "balance": 5000}`)

	req, err := http.NewRequest("POST", "http://localhost:5000/account/0000001/transfer", payload)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected status %v, but got %v", http.StatusCreated, response.Code)
	}
}
