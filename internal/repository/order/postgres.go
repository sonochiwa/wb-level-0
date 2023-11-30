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

func (p *Postgres) GetAll() ([]models.Order, error) {
	return []models.Order{}, nil
}

func (p *Postgres) GetById(orderID string) (models.Order, error) {
	return models.Order{}, nil
}

func (p *Postgres) Create(order models.Order) (string, error) {
	return "", nil
}
