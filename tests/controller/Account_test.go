package tests

import (
	"link-aja/controller/account"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func GetAllAccount(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/account", account.GetAllCustomer)

	req, err := http.NewRequest("GET", "http://localhost:5000/account", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, response.Code)
	}
}

func CreateAccount(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/account", account.GetAllCustomer)

	payload := strings.NewReader(`{"customer_name": "Bob Martin"}`)

	req, err := http.NewRequest("POST", "http://localhost:5000/account", payload)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, response.Code)
	}

	payload2 := strings.NewReader(`{"customer_name": "Linus Torvalds"}`)

	req2, err := http.NewRequest("POST", "http://localhost:5000/account", payload2)
	if err != nil {
		t.Fatal(err)
	}
	response2 := httptest.NewRecorder()

	r.ServeHTTP(response2, req2)

	if response2.Code != http.StatusCreated {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, response2.Code)
	}
}

func GetSaldoAccount(t *testing.T) {
	r := chi.NewRouter()
	r.Get("/account/{number_account}", account.GetSaldo)

	req, err := http.NewRequest("GET", "http://localhost:5000/account/0000001", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status %v, but got %v", http.StatusOK, response.Code)
	}
}
