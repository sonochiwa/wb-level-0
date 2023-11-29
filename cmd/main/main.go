package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"log"

	appConfig "github.com/sonochiwa/wb-level-0/config"
	stanClient "github.com/sonochiwa/wb-level-0/internal/clients/stan"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/server"
)

var cfg = appConfig.GetConfig()

func main() {
	log.Println("Starting API server...")

	db, err := repository.NewPostgresDB(appConfig.Postgres{
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository.NewRepository(db)

	go stanClient.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(mw.GetCors()).Handler)
	r.Get("/", server.MainPage)
	r.Mount("/", server.Routes())

	server.Run(cfg.ServerConfig.Port, r)
}
