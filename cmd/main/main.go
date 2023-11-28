package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func messageHandler(msg *stan.Msg) {
	// Обработка полученных данных, например, запись в БД и обновление кэша
	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func main() {
	// nats-streaming
	clusterID := "fddfd9c8-c093-4523-9225-567738366642"
	clientID := "client-1"
	channelName := "events"

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Публикация сообщения
	message := []byte("Hello, NATS Streaming!")
	if err := sc.Publish(channelName, message); err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	// Подписка на канал
	subscription, err := sc.Subscribe(channelName, messageHandler, stan.StartAt(pb.StartPosition_First),
		stan.DeliverAllAvailable(), stan.DurableName(clientID))
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	fmt.Printf("Subscribed to channel %s\n", channelName)

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

	// Ожидание сообщений
	select {}
}
