package order

import (
	"encoding/json"
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

	//if order.Delivery == nil {
	//	order.Delivery = &models.Delivery{}
	//}
	//
	//if order.Payment == nil {
	//	order.Payment = &models.Payment{}
	//}
	//
	//if order.Items == nil {
	//	order.Items = &[]models.OrderItems{}
	//}

	return order, err
}

// CreateOrder TODO: fix nil, nil to order.Delivery, order.Payment and add Items field
func (p *Postgres) CreateOrder(order models.Order) (string, error) {
	query := insertOrder

	delivery := &order.Delivery
	d, _ := json.Marshal(delivery)

	payment := &order.Payment
	pmnt, _ := json.Marshal(payment)

	err := p.db.QueryRow(query,
		order.TrackNumber, order.Entry, d, pmnt, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard,
	).Scan(&order.OrderUID)

	return order.OrderUID, err
}

func (p *Postgres) DeleteAllOrders() {
	query := deleteAllOrders
	p.db.Exec(query)
}
