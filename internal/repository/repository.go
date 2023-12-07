package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/models"
)

type Order interface {
	GetAllOrders() ([]models.OrderID, error)
	GetOrderById(orderID string) (models.Order, error)
	CreateOrder(items models.Order) (string, error)
	DeleteAllOrders() error
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
	}
}
