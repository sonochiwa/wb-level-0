package service

import (
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderService) GetOrderById(orderID string) (models.Order, error) {
	return s.repo.GetOrderById(orderID)
}

func (s *OrderService) CreateOrder(order models.Order) (string, error) {
	return s.repo.CreateOrder(order)
}
