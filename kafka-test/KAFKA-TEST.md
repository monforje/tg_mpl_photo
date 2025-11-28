# Kafka Test Application

Go приложение для тестирования Producer и Consumer функциональности Kafka.

## Описание

Приложение демонстрирует:
- **Producer** - отправляет сообщения в Kafka каждые 3 секунды
- **Consumer** - читает и выводит все сообщения из топика

## Структура

```
kafka-test/
├── main.go              # Точка входа
├── config.yaml          # Конфигурация
├── config/
│   └── config.go        # Загрузка конфигурации
└── kafka/
    ├── producer.go      # Producer
    └── consumer.go      # Consumer
```

## Запуск

```bash
go run main.go
```

## Конфигурация

`config.yaml`:

```yaml
kafka:
  local:
    brokers:
      - "localhost:9092"
    topic: "test-topic"
    group_id: "test-consumer-group"
  production:
    brokers:
      - "kafka-prod:9092"
    topic: "prod-topic"
    group_id: "prod-consumer-group"

env: "local"  # local или production
```

## Зависимости

```bash
go mod download
```

Основные пакеты:
- `github.com/IBM/sarama` - Kafka клиент для Go
- `gopkg.in/yaml.v3` - парсинг YAML конфигурации

## Как работает

1. **Producer** генерирует сообщения вида:
   ```
   key: "key-1"
   value: "message-1 at 2025-11-29T12:00:00Z"
   ```

2. **Consumer** читает сообщения из группы `test-consumer-group` и выводит:
   ```
   [Consumer] Message: key=key-1, value=message-1 at 2025-11-29T12:00:00Z
   ```

3. Graceful shutdown при `Ctrl+C`

## Настройка под свои нужды

- Изменить интервал отправки в `main.go` (строка с `time.NewTicker`)
- Изменить топик и brokers в `config.yaml`
- Добавить свою логику обработки сообщений в `kafka/consumer.go`