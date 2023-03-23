package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"github.com/salesof7/go_kafka-simulator-fs/infra/kafka"
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

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("Ola", "readtest", producer)

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}
	// route := route.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }
	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[1])
}
