package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Kafka KafkaConfig `yaml:"kafka"`
}

type KafkaConfig struct {
	Brokers  []string       `yaml:"brokers"`
	Producer ProducerConfig `yaml:"producer"`
}

type ProducerConfig struct {
	Topic   string        `yaml:"topic"`
	Timeout time.Duration `yaml:"timeout"`
	Retry   RetryConfig   `yaml:"retry"`
}

type RetryConfig struct {
	Max     int           `yaml:"max"`
	Backoff time.Duration `yaml:"backoff"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(".yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
