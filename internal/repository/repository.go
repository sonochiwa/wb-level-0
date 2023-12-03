package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository/order"
)

type Order interface {
	GetAllOrders() ([]models.OrderID, error)
	GetOrderById(orderID string) (models.Order, error)
	CreateOrder(order models.Order) (string, error)
	DeleteAllOrders()
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: order.NewOrderPostgres(db),
	}
}
