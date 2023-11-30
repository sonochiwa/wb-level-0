package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sonochiwa/wb-level-0/configs"
)

const (
	ordersTable     = "orders"
	orderItemsTable = "order_items"
)

func NewPostgresDB(cfg configs.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
			cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode, cfg.Host, cfg.Port))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
