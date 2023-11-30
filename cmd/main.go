package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sonochiwa/wb-level-0/internal/handlers"
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	appConfig "github.com/sonochiwa/wb-level-0/configs"
	stanClient "github.com/sonochiwa/wb-level-0/internal/clients/stan"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
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
	r.Get("/", handlers.MainPage)
	r.Mount("/", handlers.Routes())

	srv := &http.Server{
		Addr:    ":" + cfg.Postgres.Port,
		Handler: r,
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
