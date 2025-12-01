package kafka

import (
	"fmt"
	"tgbot/pkg/config"
	"tgbot/pkg/logx"
	"time"

	"github.com/IBM/sarama"
)

type Kafka struct {
	producer sarama.SyncProducer
	config   *config.KafkaConfig
}

func New(cfg *config.KafkaConfig) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()

	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true
	saramaConfig.Producer.Retry.Max = cfg.Producer.Retry.Max
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Compression = sarama.CompressionSnappy

	saramaConfig.Producer.Timeout = cfg.Producer.Timeout
	saramaConfig.Producer.Retry.Backoff = cfg.Producer.Retry.Backoff

	producer, err := sarama.NewSyncProducer(cfg.Brokers, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka producer: %w", err)
	}

	logx.Info("kafka producer initialized successfully")

	return &Kafka{
		producer: producer,
		config:   cfg,
	}, nil
}

func (k *Kafka) Producer() sarama.SyncProducer {
	return k.producer
}

func (k *Kafka) Config() *config.KafkaConfig {
	return k.config
}

func (k *Kafka) Stop() error {
	time.Sleep(100 * time.Millisecond)

	if err := k.producer.Close(); err != nil {
		return fmt.Errorf("failed to close kafka producer: %w", err)
	}

	logx.Info("kafka producer closed successfully")
	return nil
}
