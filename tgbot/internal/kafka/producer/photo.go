package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"tgbot/internal/core/event"
	"tgbot/internal/kafka"
	"tgbot/pkg/logx"
	"time"

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

func (p *PhotoProducer) Produce(ctx context.Context, event *event.PhotoUploadEvent) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := event.Validate(); err != nil {
		return fmt.Errorf("invalid photo upload event: %w", err)
	}

	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal photo upload event: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: p.kafka.Config().Producer.Topic,
		Value: sarama.ByteEncoder(data),
	}

	resultChan := make(chan error, 1)

	go func() {
		partition, offset, err := p.kafka.Producer().SendMessage(msg)
		if err != nil {
			resultChan <- fmt.Errorf("failed to send photo upload event message: %w", err)
			return
		}

		logx.Info(
			"photo upload event sent successfully",
			"partition", partition,
			"offset", offset,
			"photo_id", event.ID,
			"user_id", event.UserID,
		)
		resultChan <- nil
	}()

	select {
	case err := <-resultChan:
		if err != nil {
			return fmt.Errorf("photo upload event produce failed: %w", err)
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("photo upload event produce context done: %w", ctx.Err())
	}
}
