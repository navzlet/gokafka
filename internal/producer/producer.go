package producer

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Producer(topic string, brocker string) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brocker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("Writer started")
	defer func(writer *kafka.Writer) {
		err := writer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(writer)
	log.Println("Loop started")
	for i := 0; i < 10; i++ {
		msg := kafka.Message{
			Key:   []byte("Key"),
			Value: []byte("Привет, Kafka! " + time.Now().Format(time.RFC3339)),
		}
		err := writer.WriteMessages(context.Background(), msg)
		log.Printf("Producer produced message to topic %s\n", msg.Topic)
		if err != nil {
			log.Fatalf("Ошибка при отправке сообщения: %v", err)
		}
		log.Printf("Отправлено сообщение: %s", msg.Value)
		time.Sleep(2 * time.Second)
	}
}
