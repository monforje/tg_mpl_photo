# Kafka Test Application

Тестовое Go приложение для проверки Producer и Consumer функциональности Apache Kafka перед интеграцией с Telegram ботом.

## Назначение

Это приложение предназначено для:
- 🧪 **Тестирования** работы Kafka перед интеграцией с ботом
- 📚 **Демонстрации** базовых паттернов Producer/Consumer
- 🔍 **Отладки** конфигурации Kafka
- 📖 **Обучения** работе с Kafka в Go

## Как работает

Приложение запускает два компонента одновременно:

1. **Producer** - отправляет тестовые сообщения каждые 3 секунды:
   ```
   key: "key-1"
   value: "message-1 at 2025-11-29T12:00:00Z"
   ```

2. **Consumer** - читает все сообщения из топика в реальном времени:
   ```
   Message received: key=key-1, value=message-1 at 2025-11-29T12:00:00Z
   ```

## Структура проекта

```
kafka-test/
├── main.go              # Точка входа, запуск Producer/Consumer
├── config.yaml          # Конфигурация (brokers, topics, env)
│
├── config/
│   └── config.go        # Загрузка и парсинг конфигурации
│
└── kafka/
    ├── producer.go      # Kafka Producer
    └── consumer.go      # Kafka Consumer
```

## Быстрый старт

### 1. Убедитесь, что Kafka запущен

```bash
cd ../kafka
docker-compose up -d
```

### 2. Установите зависимости

```bash
go mod download
```

### 3. Запустите приложение

```bash
go run main.go
```

### 4. Наблюдайте за работой

Вы увидите логи:

```
2025/11/29 12:00:00 Starting Kafka test application (env: local)
2025/11/29 12:00:00 Brokers: [localhost:9092], Topic: test-topic
2025/11/29 12:00:00 Starting to read messages...
2025/11/29 12:00:00 Application is running. Press Ctrl+C to stop.
2025/11/29 12:00:03 Message sent: key=key-1, value=message-1 at 2025-11-29T12:00:03Z
2025/11/29 12:00:03 Message received: key=key-1, value=message-1 at 2025-11-29T12:00:03Z, partition=0, offset=0
2025/11/29 12:00:06 Message sent: key=key-2, value=message-2 at 2025-11-29T12:00:06Z
2025/11/29 12:00:06 Message received: key=key-2, value=message-2 at 2025-11-29T12:00:06Z, partition=0, offset=1
```

### 5. Остановка

Нажмите `Ctrl+C` для graceful shutdown:

```
2025/11/29 12:00:15 Shutting down...
2025/11/29 12:00:15 Consumer context cancelled
2025/11/29 12:00:16 Application stopped.
```

## Конфигурация

### config.yaml

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

**Параметры:**

| Параметр | Описание | Пример |
|----------|----------|--------|
| `brokers` | Адреса Kafka брокеров | `["localhost:9092"]` |
| `topic` | Имя топика для тестирования | `test-topic` |
| `group_id` | ID consumer group | `test-consumer-group` |
| `env` | Окружение (local/production) | `local` |

### Переключение окружения

Измените `env: "production"` в `config.yaml` для использования production конфигурации.

## Детали реализации

### Producer (kafka/producer.go)

```go
type Producer struct {
    writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
    writer := &kafka.Writer{
        Addr:     kafka.TCP(brokers...),
        Topic:    topic,
        Balancer: &kafka.LeastBytes{},  // Балансировка по размеру
    }
    return &Producer{writer: writer}
}

func (p *Producer) SendMessage(ctx context.Context, key, value string) error {
    msg := kafka.Message{
        Key:   []byte(key),
        Value: []byte(value),
    }
    return p.writer.WriteMessages(ctx, msg)
}
```

**Особенности:**
- Балансировщик `LeastBytes` - отправляет в партицию с наименьшим объемом данных
- Поддержка контекста для graceful shutdown
- Простая обработка ошибок

### Consumer (kafka/consumer.go)

```go
type Consumer struct {
    reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:  brokers,
        Topic:    topic,
        GroupID:  groupID,
        MinBytes: 10e3,  // 10KB
        MaxBytes: 10e6,  // 10MB
    })
    return &Consumer{reader: reader}
}

func (c *Consumer) ReadMessages(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            msg, err := c.reader.ReadMessage(ctx)
            if err != nil {
                return fmt.Errorf("failed to read message: %w", err)
            }
            log.Printf("Message received: key=%s, value=%s, partition=%d, offset=%d\n",
                string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
        }
    }
}
```

**Особенности:**
- Consumer group для координации между несколькими экземплярами
- Автоматический commit оффсетов
- MinBytes/MaxBytes для оптимизации сетевых запросов

### Главный файл (main.go)

```go
func main() {
    cfg := config.MustLoad("config.yaml")
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    producer := kafka.NewProducer(kafkaCfg.Brokers, kafkaCfg.Topic)
    consumer := kafka.NewConsumer(kafkaCfg.Brokers, kafkaCfg.Topic, kafkaCfg.GroupID)

    // Consumer в отдельной горутине
    go consumer.ReadMessages(ctx)

    // Producer с ticker
    go func() {
        ticker := time.NewTicker(3 * time.Second)
        for {
            select {
            case <-ctx.Done():
                return
            case <-ticker.C:
                producer.SendMessage(ctx, key, value)
            }
        }
    }()

    // Graceful shutdown
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
    <-sigChan
    cancel()
}
```

## Зависимости

```go
require (
    github.com/ilyakaznacheev/cleanenv v1.5.0  // Парсинг YAML
    github.com/segmentio/kafka-go v0.4.49      // Kafka клиент
)
```

**Почему segmentio/kafka-go?**
- ✅ Простой и понятный API
- ✅ Не требует CGO (в отличие от confluent-kafka-go)
- ✅ Поддержка контекстов
- ✅ Хорошая производительность

## Настройка под свои нужды

### Изменить интервал отправки

В `main.go`:

```go
ticker := time.NewTicker(5 * time.Second)  // Было 3 секунды
```

### Изменить формат сообщений

В `main.go`:

```go
// Было
value := fmt.Sprintf("message-%d at %s", counter, time.Now().Format(time.RFC3339))

// Стало (JSON)
type Message struct {
    ID        int       `json:"id"`
    Content   string    `json:"content"`
    Timestamp time.Time `json:"timestamp"`
}
msg := Message{ID: counter, Content: "Hello", Timestamp: time.Now()}
value, _ := json.Marshal(msg)
```

### Добавить обработку сообщений

В `kafka/consumer.go`:

```go
func (c *Consumer) ReadMessages(ctx context.Context) error {
    for {
        msg, err := c.reader.ReadMessage(ctx)
        if err != nil {
            return err
        }

        // Добавить свою логику
        if err := processMessage(msg); err != nil {
            log.Printf("Processing error: %v\n", err)
        }
    }
}

func processMessage(msg kafka.Message) error {
    // Ваша бизнес-логика
    var data MyDataType
    json.Unmarshal(msg.Value, &data)
    // ...
    return nil
}
```

### Добавить несколько топиков

```go
topics := []string{"topic-1", "topic-2", "topic-3"}

for _, topic := range topics {
    producer := kafka.NewProducer(brokers, topic)
    go producer.SendMessage(ctx, "key", "value")
}
```

## Интеграция с Telegram ботом

После успешного тестирования можно интегрировать Kafka в бота:

### 1. Скопировать Producer/Consumer

```bash
cp kafka-test/kafka/* tgbot/internal/kafka/
```

### 2. Отправлять события из бота

```go
// В RegService после создания пользователя
func (r *RegService) Reg(tgID int64, username string) error {
    // ... создание пользователя ...

    // Отправить событие в Kafka
    event := map[string]interface{}{
        "event_type": "user.registered",
        "user_id":    id.String(),
        "tg_id":      tgID,
        "username":   username,
        "timestamp":  time.Now().Unix(),
    }
    eventJSON, _ := json.Marshal(event)
    r.kafkaProducer.SendMessage(ctx, id.String(), string(eventJSON))

    return nil
}
```

### 3. Обрабатывать команды через Kafka

```go
// Сервис для обработки команд из Kafka
type CommandProcessor struct {
    consumer *kafka.Consumer
    bot      *tele.Bot
}

func (p *CommandProcessor) ProcessCommands(ctx context.Context) {
    for {
        msg, _ := p.consumer.ReadMessage(ctx)
        
        var cmd Command
        json.Unmarshal(msg.Value, &cmd)
        
        // Выполнить команду через бота
        p.bot.Send(cmd.ChatID, cmd.Response)
    }
}
```

## Тестирование

### Проверка производительности

```go
// Добавить в main.go
ticker := time.NewTicker(100 * time.Millisecond)  // 10 сообщений/сек

var sentCount, receivedCount atomic.Int64

go func() {
    for range ticker.C {
        producer.SendMessage(ctx, key, value)
        sentCount.Add(1)
    }
}()

go func() {
    for {
        consumer.ReadMessage(ctx)
        receivedCount.Add(1)
    }
}()

// Каждые 10 секунд выводить статистику
go func() {
    for range time.Tick(10 * time.Second) {
        log.Printf("Sent: %d, Received: %d\n", sentCount.Load(), receivedCount.Load())
    }
}()
```

### Проверка устойчивости

```bash
# Запустить приложение
go run main.go

# В другом терминале остановить Kafka
docker-compose -f ../kafka/docker-compose.yml down

# Подождать 10 секунд

# Запустить снова
docker-compose -f ../kafka/docker-compose.yml up -d

# Приложение должно автоматически переподключиться
```

## Troubleshooting

### "Connection refused" при запуске

```bash
# Проверить, запущен ли Kafka
docker ps | grep kafka

# Проверить порт
telnet localhost 9092

# Если не запущен - запустить
cd ../kafka && docker-compose up -d
```

### Сообщения не доставляются

```bash
# Проверить топик
docker exec -it kafka kafka-topics.sh --list --bootstrap-server localhost:9092

# Создать топик вручную, если нужно
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --partitions 1 \
  --replication-factor 1
```

### Consumer не читает старые сообщения

Измените в `kafka/consumer.go`:

```go
reader := kafka.NewReader(kafka.ReaderConfig{
    Brokers:  brokers,
    Topic:    topic,
    GroupID:  groupID,
    MinBytes: 10e3,
    MaxBytes: 10e6,
    StartOffset: kafka.FirstOffset,  // Добавить эту строку
})
```

## Следующие шаги

После успешного тестирования:

1. ✅ Убедиться, что Producer и Consumer работают
2. ✅ Понять паттерны отправки/получения сообщений
3. 🔄 Скопировать код в проект с ботом
4. 🔄 Интегрировать Producer для событий бота
5. 🔄 Добавить Consumer для обработки команд
6. 📊 Настроить мониторинг и алертинг

## Полезные команды

```bash
# Запустить с verbose логами
go run main.go 2>&1 | tee kafka-test.log

# Проверить количество сообщений в топике
docker exec -it kafka kafka-run-class.sh kafka.tools.GetOffsetShell \
  --broker-list localhost:9092 \
  --topic test-topic

# Прочитать все сообщения из топика
docker exec -it kafka kafka-console-consumer.sh \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --from-beginning
```