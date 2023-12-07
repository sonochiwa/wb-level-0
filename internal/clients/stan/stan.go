package stan

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
	config "github.com/sonochiwa/wb-level-0/configs"
	mc "github.com/sonochiwa/wb-level-0/internal/memcache"
	"github.com/sonochiwa/wb-level-0/internal/models"
)

var cfg = config.GetConfig()
var sc stan.Conn

func messageHandler(msg *stan.Msg) {
	data := string(msg.Data)
	err := saveToDB(data)
	if err != nil {
		log.Printf("Error processing data: %v", err)
	}

	//err = saveToMemCache()
	//if err != nil {
	//	log.Printf(err)
	//}
}

func saveToMemCache(data string) (interface{}, bool) {
	cache := mc.New(5*time.Minute, 10*time.Minute)
	cache.Set("myKey", "My value", 5*time.Minute)
	i, found := cache.Get("myKey")
	if !found {
		return nil, found
	}
	return i, false
}

func saveToDB(data string) error {
	order := models.Order{}

	err := json.Unmarshal([]byte(data), &order)
	if err != nil {
		return err
	}

	//repository.NewRepository().CreateOrder(order)

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
	subscription, err := sc.Subscribe(cfg.Stan.ChannelName, messageHandler, stan.StartAt(pb.StartPosition_First),
		stan.DeliverAllAvailable())
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
