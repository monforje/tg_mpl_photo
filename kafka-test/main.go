package main

import (
	"context"
	"fmt"
	"kafka-test/config"
	"kafka-test/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad("config.yaml")
	kafkaCfg := cfg.GetKafkaConfig()

	log.Printf("Starting Kafka test application (env: %s)\n", cfg.Env)
	log.Printf("Brokers: %v, Topic: %s\n", kafkaCfg.Brokers, kafkaCfg.Topic)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	producer := kafka.NewProducer(kafkaCfg.Brokers, kafkaCfg.Topic)
	defer producer.Close()

	consumer := kafka.NewConsumer(kafkaCfg.Brokers, kafkaCfg.Topic, kafkaCfg.GroupID)
	defer consumer.Close()

	go func() {
		if err := consumer.ReadMessages(ctx); err != nil && err != context.Canceled {
			log.Printf("Consumer error: %v\n", err)
		}
	}()

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				counter++
				key := fmt.Sprintf("key-%d", counter)
				value := fmt.Sprintf("message-%d at %s", counter, time.Now().Format(time.RFC3339))

				if err := producer.SendMessage(ctx, key, value); err != nil {
					log.Printf("Producer error: %v\n", err)
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log.Println("Application is running. Press Ctrl+C to stop.")
	<-sigChan

	log.Println("Shutting down...")
	cancel()
	time.Sleep(1 * time.Second)
	log.Println("Application stopped.")
}
