package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sonochiwa/wb-level-0/configs"
	sc "github.com/sonochiwa/wb-level-0/internal/clients/stan"
	"github.com/sonochiwa/wb-level-0/internal/handler"
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"github.com/sonochiwa/wb-level-0/internal/service"
)

var cfg = configs.GetConfig()

func main() {
	log.Println("Starting API server...")

	db, err := repository.NewPostgresDB(configs.Postgres{
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

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go sc.New()

	srv := &http.Server{
		Addr:    ":" + cfg.Postgres.Port,
		Handler: handlers.InitRoutes(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
