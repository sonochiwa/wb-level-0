package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) getAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.services.Order.GetAllOrders()
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) getOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "orderID")

	order, err := h.services.Order.GetOrderById(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *Handler) createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	order, err := h.services.Order.CreateOrder()
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *Handler) deleteAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	h.services.Order.DeleteAllOrders()

	json.NewEncoder(w).Encode("ok")
}
