package stan

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"

	config "github.com/sonochiwa/wb-level-0/configs"
)

var cfg = config.GetConfig()
var sc stan.Conn

func messageHandler(msg *stan.Msg) {
	data := string(msg.Data)
	err := saveToDB(data)
	if err != nil {
		log.Printf("Error processing data: %v", err)
	}

	fmt.Printf("Received a message: %s\n", string(msg.Data))
}

func saveToDB(data string) error {
	return nil
}

func PublishMessage(publishChan <-chan string) error {
	message := <-publishChan
	if err := sc.Publish(cfg.Stan.ChannelName, []byte(message)); err != nil {
		log.Printf("Error publishing message: %v", err)
		return err
	}
	log.Printf("Published message to channel %s: %s\n", message)
	return nil
}

func New() {
	sc, err := stan.Connect(cfg.Stan.ClusterID, cfg.Stan.ClientID, stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	subscription, err := sc.Subscribe(cfg.Stan.ChannelName, messageHandler, stan.StartAt(pb.StartPosition_First),
		stan.DeliverAllAvailable(), stan.DurableName(cfg.Stan.ClientID))
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
