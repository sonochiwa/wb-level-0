package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository/order"
)

type Order interface {
	GetAll() ([]models.Order, error)
	GetById(orderID string) (models.Order, error)
	Create(order models.Order) (string, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: order.NewOrderPostgres(db),
	}
}
