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

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetById(orderID string) (models.Order, error) {
	return s.repo.GetById(orderID)
}

func (s *OrderService) Create(order models.Order) (string, error) {
	return s.repo.Create(order)
}
