package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sonochiwa/wb-level-0/configs"
)

var cfg = configs.GetConfig()
var DB *sqlx.DB

func init() {
	db, err := GetDB(DB)
	if err != nil {
		return
	}
	DB = db
}

func GetDB(DB *sqlx.DB) (*sqlx.DB, error) {
	db, err := NewPostgresDB(configs.Postgres{
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		return nil, err
	}
	DB = db
	return DB, nil
}

func NewPostgresDB(cfg configs.Postgres) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres",
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
