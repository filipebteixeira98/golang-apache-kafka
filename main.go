package main

import (
	"fmt"
	"log"

	// "github.com/filipebteixeira98/go-kafka/app/routes"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	kafkap "github.com/filipebteixeira98/go-kafka/app/kafka"
	"github.com/filipebteixeira98/go-kafka/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkap.Produce(msg)
	}

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("example", "readtest", producer)
	// for {
	// 	_ = 1
	// }

	// r := routes.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// r.LoadPositions()
	// rJSON, _ := r.ExportJSONPositions()
	// fmt.Println(rJSON[0])
}
