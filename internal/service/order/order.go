package order

import (
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"github.com/sonochiwa/wb-level-0/internal/repository/order"
)

type Service struct {
	repo repository.Repository
}

func NewOrderService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrder(order order.Order) (string, error) {
	return s.repo.Create(order)
}
