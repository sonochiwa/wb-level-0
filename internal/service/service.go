package service

import (
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
)

type Order interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderById(orderID string) (models.Order, error)
	CreateOrder(order models.Order) (string, error)
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
