package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/repository/order"
)

type Order interface {
	GetIdentifiers()
	GetById()
	Create(order order.Order) (string, error)
}

type Repository struct {
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Order: NewRepository(db),
	}
}
