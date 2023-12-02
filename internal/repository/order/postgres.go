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
	var orders []models.Order

	query := selectOrders
	err := p.db.Select(&orders, query)

	for _, order := range orders {
		if order.Delivery == nil {
			order.Delivery = &models.Delivery{}
		}

		if order.Payment == nil {
			order.Payment = &models.Payment{}
		}

		if order.Items == nil {
			order.Items = &[]models.OrderItems{}
		}
	}

	return orders, err
}

func (p *Postgres) GetOrderById(orderID string) (models.Order, error) {
	var order models.Order

	query := selectOrderByID
	err := p.db.Get(&order, query, orderID)

	if order.Delivery == nil {
		order.Delivery = &models.Delivery{}
	}

	if order.Payment == nil {
		order.Payment = &models.Payment{}
	}

	if order.Items == nil {
		order.Items = &[]models.OrderItems{}
	}

	return order, err
}

func (p *Postgres) CreateOrder(order models.Order) (string, error) {
	query := createOrder
	err := p.db.QueryRow(query, order.TrackNumber).Scan(&order.OrderUID)

	return order.OrderUID, err
}
