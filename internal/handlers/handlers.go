package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}

func getOrderIdentifiers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hello world"))
}

func getOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "orderID")
	fmt.Println(id)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("ok"))
}

func Routes() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/orders/identifiers", getOrderIdentifiers)
		r.Get("/orders/{orderID}", getOrderByID)
		r.Post("/orders", createOrder)
	})
	return rg
}
