package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	GroupID string   `yaml:"group_id"`
}

type Config struct {
	Kafka struct {
		Local      KafkaConfig `yaml:"local"`
		Production KafkaConfig `yaml:"production"`
	} `yaml:"kafka"`
	Env string `yaml:"env"`
}

func MustLoad(configPath string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}

func (c *Config) GetKafkaConfig() KafkaConfig {
	if c.Env == "production" {
		return c.Kafka.Production
	}
	return c.Kafka.Local
}
