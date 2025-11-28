# Kafka Test Application

Тестовое приложение для работы с Apache Kafka на Go.

## Структура проекта

- **kafka/** - Kafka сервер в Docker ([подробнее](kafka/KAFKA.md))
- **kafka-test/** - Go приложение для тестирования ([подробнее](kafka-test/KAFKA-TEST.md))

## Быстрый старт

1. Запустить Kafka сервер:
   ```bash
   cd kafka
   docker-compose up -d
   ```

2. Запустить тестовое приложение:
   ```bash
   cd kafka-test
   go run main.go
   ```

3. Остановить всё:
   ```bash
   Ctrl+C  # остановить приложение
   cd kafka && docker-compose down  # остановить Kafka
   ```

## Требования

- Docker
- Go 1.21+