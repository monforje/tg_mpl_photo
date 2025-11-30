package kafka

import (
	"tgbot/pkg/config"
	"tgbot/pkg/logx"

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

	producer, err := sarama.NewSyncProducer(cfg.Brokers, saramaConfig)
	if err != nil {
		return nil, err
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
	if err := k.producer.Close(); err != nil {
		return err
	}
	logx.Info("kafka producer closed")
	return nil
}
