package producer

import (
	"encoding/json"
	"tgbot/internal/core/event"
	"tgbot/internal/kafka"
	"tgbot/pkg/logx"

	"github.com/IBM/sarama"
)

type PhotoUploadEventProducer struct {
	kafka *kafka.Kafka
}

func NewPhotoUploadEventProducer(kafka *kafka.Kafka) *PhotoUploadEventProducer {
	return &PhotoUploadEventProducer{
		kafka: kafka,
	}
}

func (p *PhotoUploadEventProducer) Produce(event *event.PhotoUploadEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: p.kafka.Config().Producer.Topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := p.kafka.Producer().SendMessage(msg)
	if err != nil {
		return err
	}

	logx.Info(
		"photo upload event sent successfully",
		"partition", partition,
		"offset", offset,
		"photo_id", event.ID,
	)

	return nil
}
