package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load(".env")

	rabbitConnString := os.Getenv("AMQP_URL")
	if rabbitConnString == "" {
		log.Fatalf("AQMP_URL environment is not set")
	}

	connection, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Couldn't connect to RabbitMQ: %v", err)
	}
	defer connection.Close()
	fmt.Println("Peril game server connected to RabbitMQ server successfully")

	publishCh, err := connection.Channel()
	if err != nil {
		log.Fatalf("Couldn't create channel: %v", err)
	}
	err = pubsub.PublishJSON(
		publishCh,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)

	if err != nil {
		log.Printf("Couldn't publish time: %v", err)
	}
	fmt.Println("Pause message sent!")
}
