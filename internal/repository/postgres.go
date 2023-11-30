package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	appConfig "github.com/sonochiwa/wb-level-0/configs"
)

const (
	ordersTable     = "orders"
	orderItemsTable = "order_items"
)

func NewPostgresDB(cfg appConfig.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
