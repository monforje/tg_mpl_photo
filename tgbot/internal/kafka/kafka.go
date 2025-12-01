package kafka

import (
	"fmt"
	"tgbot/pkg/config"
	"tgbot/pkg/logx"

	"github.com/IBM/sarama"
)

type Kafka struct {
	producer sarama.AsyncProducer
	config   *config.KafkaConfig
}

func New(cfg *config.KafkaConfig) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()

	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true
	saramaConfig.Producer.Retry.Max = cfg.Producer.Retry.Max
	saramaConfig.Producer.RequiredAcks = sarama.WaitForLocal
	saramaConfig.Producer.Compression = sarama.CompressionSnappy
	saramaConfig.Producer.Timeout = cfg.Producer.Timeout
	saramaConfig.Producer.Retry.Backoff = cfg.Producer.Retry.Backoff
	saramaConfig.Producer.Idempotent = true
	saramaConfig.Net.MaxOpenRequests = 1

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, saramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka producer: %w", err)
	}

	k := &Kafka{
		producer: producer,
		config:   cfg,
	}

	go k.handleResults()

	logx.Info("kafka async producer initialized successfully")

	return k, nil
}

func (k *Kafka) handleResults() {
	for {
		select {
		case msg, ok := <-k.producer.Successes():
			if !ok {
				return
			}
			logx.Info(
				"message sent successfully",
				"topic", msg.Topic,
				"partition", msg.Partition,
				"offset", msg.Offset,
			)
		case err, ok := <-k.producer.Errors():
			if !ok {
				return
			}
			logx.Error(
				"failed to send message",
				"error", err.Err,
				"topic", err.Msg.Topic,
			)
		}
	}
}

func (k *Kafka) Producer() sarama.AsyncProducer {
	return k.producer
}

func (k *Kafka) Config() *config.KafkaConfig {
	return k.config
}

func (k *Kafka) Stop() error {
	if err := k.producer.Close(); err != nil {
		return fmt.Errorf("failed to close kafka producer: %w", err)
	}

	logx.Info("kafka producer closed successfully")
	return nil
}
