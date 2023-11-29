package stan

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"

	appConfig "github.com/sonochiwa/wb-level-0/config"
)

var cfg = appConfig.GetConfig()

func messageHandler(msg *stan.Msg) {
	// Обработка полученных данных, например, запись в БД и обновление кэша
	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func New() {
	sc, err := stan.Connect(cfg.Stan.ClusterID, cfg.Stan.ClientID,
		stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Публикация сообщения
	message := []byte("Hello, NATS Streaming!")
	if err := sc.Publish(cfg.Stan.ChannelName, message); err != nil {
		log.Fatalf("Error publishing message: %v", err)
	}

	// Подписка на канал
	subscription, err := sc.Subscribe(cfg.Stan.ChannelName, messageHandler, stan.StartAt(pb.StartPosition_First),
		stan.DeliverAllAvailable(), stan.DurableName(cfg.Stan.ClientID))
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	fmt.Printf("Subscribed to channel %s\n", cfg.Stan.ChannelName)

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
}
