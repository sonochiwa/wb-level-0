package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"net/http"
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
	var t models.Order
	json.NewDecoder(r.Body).Decode(&t)

	order, err := h.services.Order.CreateOrder(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(order)
}
