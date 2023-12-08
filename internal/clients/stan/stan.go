package stan

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	config "github.com/sonochiwa/wb-level-0/configs"
	mc "github.com/sonochiwa/wb-level-0/internal/memcache"
	"github.com/sonochiwa/wb-level-0/internal/models"
	"github.com/sonochiwa/wb-level-0/internal/repository"
)

var cfg = config.GetConfig()
var sc stan.Conn

func messageHandler(msg *stan.Msg) {
	data := string(msg.Data)
	err := saveToDB(data)
	if err != nil {
		log.Printf("Error processing data: %v", err)
		return
	}

	err = saveToCache(data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveToCache(data string) error {
	order := models.Order{}

	err := json.Unmarshal([]byte(data), &order)
	if err != nil {
		return err
	}

	if err := mc.MC.Set(&cache.Item{
		Ctx: context.TODO(), Key: order.OrderUID, Value: data, TTL: time.Hour}); err != nil {
		return err
	}

	return nil
}

// TODO: unification repo interface
func saveToDB(data string) error {
	order := models.Order{}
	p := &repository.Postgres{DB: repository.DB}

	err := json.Unmarshal([]byte(data), &order)

	if err != nil {
		return err
	}

	err = p.CreateOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func PublishMessage(message []byte) error {
	err := GetStanConnection().Publish(cfg.Stan.ChannelName, message)
	if err != nil {
		return err
	}

	return nil
}

func GetStanConnection() stan.Conn {
	return sc
}

func New(clusterID, clientID, natsURL string) (stan.Conn, error) {
	natsConn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	sc, err = stan.Connect(clusterID, clientID, stan.NatsConn(natsConn))
	if err != nil {
		return nil, err
	}

	return sc, nil
}

func Subscribe(sc stan.Conn) {
	subscription, err := sc.Subscribe(cfg.Stan.ChannelName, messageHandler)
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	}
	fmt.Printf("Subscribed to channel %s\n", cfg.Stan.ChannelName)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	err = subscription.Unsubscribe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Unsubscribed from channel")

	select {}
}
