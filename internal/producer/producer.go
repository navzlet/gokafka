package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
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
			Key:   []byte("ZXC"),
			Value: []byte("Привет, KaKafka!" + string(rune(i))),
		}
		err := writer.WriteMessages(context.Background(), msg)
		log.Printf("Producer produced message to topic %s\n", msg.Topic)
		if err != nil {
			log.Fatalf("Ошибка при отправке сообщения: %v", err)
		}
		log.Printf("Отправлено сообщение: %s", msg.Value)
	}
}
