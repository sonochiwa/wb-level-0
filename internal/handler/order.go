package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func getAllOrders(w http.ResponseWriter, r *http.Request) {
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
