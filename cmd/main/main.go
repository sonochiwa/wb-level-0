package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	mw "github.com/sonochiwa/wb-level-0/internal/middleware"
	"github.com/sonochiwa/wb-level-0/internal/server"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func messageHandler(msg *stan.Msg) {
	// Обработка полученных данных, например, запись в БД и обновление кэша
	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func main() {
	log.Println("Starting API server...")

	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	go func() {
		sc, err := stan.Connect(viper.GetString("stan.ClusterID"), viper.GetString("stan.ClientID"),
			stan.NatsURL(stan.DefaultNatsURL))
		if err != nil {
			log.Fatal(err)
		}
		defer sc.Close()

		// Публикация сообщения
		message := []byte("Hello, NATS Streaming!")
		if err := sc.Publish(viper.GetString("stan.ChannelName"), message); err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}

		// Подписка на канал
		subscription, err := sc.Subscribe(viper.GetString("stan.ChannelName"), messageHandler, stan.StartAt(pb.StartPosition_First),
			stan.DeliverAllAvailable(), stan.DurableName(viper.GetString("stan.ClientID")))
		if err != nil {
			log.Fatalf("Error subscribing to channel: %v", err)
		}
		fmt.Printf("Subscribed to channel %s\n", viper.GetString("stan.ChannelName"))

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
		select {}
	}()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.New(mw.GetCors()).Handler)
	r.Get("/", server.MainPage)
	r.Mount("/", server.Routes())

	http.ListenAndServe(":"+viper.GetString("server.port"), r)
}
