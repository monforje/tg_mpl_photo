package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"tgbot/internal/core/event"
	"tgbot/internal/kafka"

	"github.com/IBM/sarama"
)

type PhotoProducer struct {
	kafka *kafka.Kafka
}

func NewPhotoProducer(kafka *kafka.Kafka) *PhotoProducer {
	return &PhotoProducer{
		kafka: kafka,
	}
}

func (p *PhotoProducer) PublishPhoto(ctx context.Context, e *event.PhotoUploadEvent) error {
	if err := e.Validate(); err != nil {
		return fmt.Errorf("invalid photo upload event: %w", err)
	}

	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal photo upload event: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.kafka.Config().Producer.Topic,
		Key:   sarama.StringEncoder(e.UserID.String()),
		Value: sarama.ByteEncoder(data),
	}

	select {
	case p.kafka.Producer().Input() <- msg:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("photo upload event produce context done: %w", ctx.Err())
	}
}
