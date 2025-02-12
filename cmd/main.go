package main

import (
	"github.com/segmentio/kafka-go"
	"gokafka/internal/consumer"
	"gokafka/internal/producer"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	topic := "test123"
	broker := os.Getenv("KAFKA_BROKER")
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(2)

	// Запуск продюсера в отдельной горутине
	go func() {
		defer wg.Done()
		producer.Producer(topic, broker)
	}()

	// Запуск консьюмера в отдельной горутине
	go func() {
		defer wg.Done()
		consumer.Consumer(topic, broker)
	}()

	wg.Wait() // Блокировка главной горутины до завершения работы
}
