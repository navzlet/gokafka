package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func Consumer(topic string, brocker string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brocker},
		Topic:   topic,
		GroupID: "example-group",
	})

	defer func(reader *kafka.Reader) {
		err := reader.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(reader)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Ошибка при чтении сообщения: %v", err)
		}
		log.Printf("Получено сообщение: key = %s, value = %s, offset = %d",
			string(msg.Key), string(msg.Value), msg.Offset)
	}
}
