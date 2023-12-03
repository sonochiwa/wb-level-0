package service

import (
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
)

type Order interface {
	GetAllOrders() ([]models.OrderID, error)
	GetOrderById(orderID string) (models.Order, error)
	CreateOrder(order models.Order) (string, error)
	DeleteAllOrders()
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
