package main

import (
	"fmt"
	"link-aja/controller/account"
	"link-aja/controller/transaksi"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		// Check All Account
		r.Get("/account", account.GetAllCustomer)
		// Check Saldo
		r.Get("/account/{number_account}", account.GetSaldo)
		// Create Account
		r.Post("/account", account.CreateAccount)
		// Transfer
		r.Post("/account/{number_account}/transfer", transaksi.Transfer)
		// Add Balance
		r.Post("/account/{number_account}/top-up", transaksi.AddBalance)
	})

	fmt.Println("http://localhost:5000")

	http.ListenAndServe(":5000", r)
}
