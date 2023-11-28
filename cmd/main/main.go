package main

import (
	"context"
	"fmt"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"

	appConfig "github.com/sonochiwa/wb-level-0/config"
	mware "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/server"
)

var cfg = appConfig.GetConfig()

func messageHandler(msg *stan.Msg) {
	// Обработка полученных данных, например, запись в БД и обновление кэша
	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func main() {
	// nats-streaming
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(cfg.NatsStreaming.ClusterID, cfg.NatsStreaming.ClientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Публикация сообщения
	message := []byte("Hello, NATS Streaming!")
	if err := sc.Publish(cfg.NatsStreaming.ChannelName, message); err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	// Подписка на канал
	subscription, err := sc.Subscribe(cfg.NatsStreaming.ChannelName, messageHandler, stan.StartAt(pb.StartPosition_First),
		stan.DeliverAllAvailable(), stan.DurableName(cfg.NatsStreaming.ClientID))
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	fmt.Printf("Subscribed to channel %s\n", cfg.NatsStreaming.ChannelName)

	// Ждем сигнала завершения (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	// Отписываемся от канала при завершении
	err = subscription.Unsubscribe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unsubscribed from channel")
	//Ожидание сообщений
	//select {}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(mware.GetCors()).Handler)
	r.Get("/", server.MainPage)
	r.Mount("/", server.Routes())

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.Server.IdleTimeout) * time.Second,
	}

	go func() {
		log.Println("Already available")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	<-ch
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server was stopped")
}
