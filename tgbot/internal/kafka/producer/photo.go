package producer

import (
	"context"
	"encoding/json"
	"tgbot/internal/core/event"
	"tgbot/internal/kafka"
	"tgbot/pkg/logx"
	"time"

	"github.com/IBM/sarama"
)

const (
	defaultProduceTimeout = 10 * time.Second
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
	ctx, cancel := context.WithTimeout(ctx, defaultProduceTimeout)
	defer cancel()

	if err := event.Validate(); err != nil {
		return err
	}

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.kafka.Config().Producer.Topic,
		Value: sarama.ByteEncoder(data),
	}

	resultChan := make(chan error, 1)

	go func() {
		partition, offset, err := p.kafka.Producer().SendMessage(msg)
		if err != nil {
			resultChan <- err
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
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
