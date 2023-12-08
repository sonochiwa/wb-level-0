package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/sonochiwa/wb-level-0/configs"
	sc "github.com/sonochiwa/wb-level-0/internal/clients/stan"
	"github.com/sonochiwa/wb-level-0/internal/handler"
	"github.com/sonochiwa/wb-level-0/internal/repository"
	"github.com/sonochiwa/wb-level-0/internal/service"
)

var cfg = configs.GetConfig()

func main() {
	log.Println("Starting API server...")
	db, err := repository.GetDB(repository.DB)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.New(db)
	services := service.New(repos)
	handlers := handler.New(services)

	stanConn, err := sc.New(cfg.Stan.ClusterID, cfg.Stan.ClientID, stan.DefaultNatsURL)
	if err != nil {
		panic(err)
	}

	go sc.Subscribe(stanConn)

	defer stanConn.Close()

	srv := &http.Server{
		Addr:    ":" + cfg.ServerConfig.Port,
		Handler: handlers.InitRoutes(),
	}

	fmt.Printf("Server running on http://%s:%s\n", cfg.ServerConfig.Host, cfg.ServerConfig.Port)

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
