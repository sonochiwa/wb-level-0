package order

import (
	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/models"
)

type Postgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllOrders() ([]models.Order, error) {
	return []models.Order{}, nil
}

func (p *Postgres) GetOrderById(orderID string) (models.Order, error) {
	return models.Order{}, nil
}

func (p *Postgres) CreateOrder(order models.Order) (string, error) {
	return "", nil
}
