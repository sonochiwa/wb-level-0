package repository

import (
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/sonochiwa/wb-level-0/internal/models"
)

type Postgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) GetAllOrders() ([]models.OrderID, error) {
	var orders []models.OrderID

	query := selectOrders
	err := p.db.Select(&orders, query)

	return orders, err
}

func (p *Postgres) GetOrderById(orderID string) (models.Order, error) {
	var order models.Order

	query := selectOrderByID
	err := p.db.Get(&order, query, orderID)

	return order, err
}

func (p *Postgres) CreateOrder(order models.Order) (string, error) {
	query := insertOrder

	delivery, err := json.Marshal(&order.Delivery)
	if err != nil {
		log.Fatal(err)
	}

	payment, err := json.Marshal(&order.Payment)
	if err != nil {
		log.Fatal(err)
	}

	items, err := json.Marshal(&order.Items)
	if err != nil {
		log.Fatal(err)
	}

	err = p.db.QueryRow(query,
		order.TrackNumber, order.Entry, delivery, payment, items, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard,
	).Scan(&order.OrderUID)

	return order.OrderUID, err
}

func (p *Postgres) DeleteAllOrders() error {
	query := deleteAllOrders
	_, err := p.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
