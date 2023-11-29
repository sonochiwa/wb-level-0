package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	appConfig "github.com/sonochiwa/wb-level-0/config"
	stanClient "github.com/sonochiwa/wb-level-0/internal/clients/stan"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/server"
)

var cfg = appConfig.GetConfig()

func main() {
	log.Println("Starting API server...")

	go stanClient.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(mw.GetCors()).Handler)
	r.Get("/", server.MainPage)
	r.Mount("/", server.Routes())

	// gracefully shutdown
	srv := &http.Server{Addr: ":" + cfg.ServerConfig.Port, Handler: r}
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
