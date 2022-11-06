package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
	"simulater-uber-car/application"
	"simulater-uber-car/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

// {"clientId":"1", "routeId":"1"}
// {"clientId":"2", "routeId":"2"}
// {"clientId":"3", "routeId":"3"}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go application.Produce(msg)
	}
}
