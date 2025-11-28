package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	return &Consumer{reader: reader}
}

func (c *Consumer) ReadMessages(ctx context.Context) error {
	log.Println("Starting to read messages...")

	for {
		select {
		case <-ctx.Done():
			log.Println("Consumer context cancelled")
			return ctx.Err()
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				return fmt.Errorf("failed to read message: %w", err)
			}

			log.Printf("Message received: key=%s, value=%s, partition=%d, offset=%d\n",
				string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}
