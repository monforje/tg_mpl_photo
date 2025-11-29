# Kafka Server

Apache Kafka брокер в Docker контейнере с использованием современного KRaft режима (без Zookeeper).

## Особенности

- 🚀 **KRaft режим** - современная архитектура без Zookeeper
- 📦 **Bitnami Kafka 4.0.0** - стабильная версия
- 🔄 **Персистентное хранилище** - данные сохраняются между перезапусками
- 🎯 **Готов к интеграции** - с Telegram ботом и другими сервисами

## Быстрый старт

### Запуск

```bash
docker-compose up -d
```

Проверка запуска:

```bash
docker logs -f kafka
```

Успешный запуск выглядит так:

```
[2025-11-29 12:00:00] INFO Kafka Server started
[2025-11-29 12:00:01] INFO KRaft mode enabled
[2025-11-29 12:00:02] INFO Listening on PLAINTEXT://localhost:9092
```

### Остановка

```bash
# Остановить с сохранением данных
docker-compose down

# Остановить и удалить данные
docker-compose down -v
```

## Конфигурация

### Основные параметры

| Параметр | Значение | Описание |
|----------|----------|----------|
| **Image** | `bitnamilegacy/kafka:4.0.0` | Стабильная версия Kafka |
| **Режим** | KRaft | Без Zookeeper |
| **Node ID** | 1 | Идентификатор узла |
| **Listeners** | PLAINTEXT:9092, CONTROLLER:9093 | Порты для клиентов и контроллера |

### Порты

- **9092** - PLAINTEXT listener (для клиентов, producer/consumer)
- **9093** - CONTROLLER listener (внутренний, для координации)

### Volumes

Данные сохраняются в `./volumes/`:

```
kafka/volumes/
├── data/           # Логи и партиции топиков
└── config/         # Конфигурационные файлы
```

> ⚠️ **Важно:** Директория `volumes/` добавлена в `.gitignore` для предотвращения коммита данных

## Работа с топиками

### Создать топик

```bash
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic bot-events \
  --partitions 3 \
  --replication-factor 1
```

**Параметры:**
- `bot-events` - имя топика
- `partitions 3` - количество партиций (для параллелизма)
- `replication-factor 1` - фактор репликации (для одного брокера всегда 1)

### Список топиков

```bash
docker exec -it kafka kafka-topics.sh --list \
  --bootstrap-server localhost:9092
```

### Описание топика

```bash
docker exec -it kafka kafka-topics.sh --describe \
  --bootstrap-server localhost:9092 \
  --topic bot-events
```

Вывод:

```
Topic: bot-events       PartitionCount: 3       ReplicationFactor: 1
    Topic: bot-events   Partition: 0    Leader: 1       Replicas: 1     Isr: 1
    Topic: bot-events   Partition: 1    Leader: 1       Replicas: 1     Isr: 1
    Topic: bot-events   Partition: 2    Leader: 1       Replicas: 1     Isr: 1
```

### Изменить конфигурацию топика

```bash
docker exec -it kafka kafka-configs.sh --alter \
  --bootstrap-server localhost:9092 \
  --entity-type topics \
  --entity-name bot-events \
  --add-config retention.ms=604800000
```

### Удалить топик

```bash
docker exec -it kafka kafka-topics.sh --delete \
  --bootstrap-server localhost:9092 \
  --topic bot-events
```

## Тестирование

### Console Producer

Отправка сообщений вручную:

```bash
docker exec -it kafka kafka-console-producer.sh \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --property "parse.key=true" \
  --property "key.separator=:"
```

Затем вводите сообщения в формате `key:value`:

```
user1:Hello from console producer!
user2:Another message
```

### Console Consumer

Чтение сообщений:

```bash
# С начала топика
docker exec -it kafka kafka-console-consumer.sh \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --from-beginning \
  --property print.key=true \
  --property key.separator=":"

# Только новые сообщения
docker exec -it kafka kafka-console-consumer.sh \
  --bootstrap-server localhost:9092 \
  --topic test-topic
```

### Группы потребителей

Просмотр consumer groups:

```bash
docker exec -it kafka kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --list
```

Детали группы:

```bash
docker exec -it kafka kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group test-consumer-group \
  --describe
```

Сброс оффсетов:

```bash
docker exec -it kafka kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group test-consumer-group \
  --topic test-topic \
  --reset-offsets \
  --to-earliest \
  --execute
```

## Интеграция с Telegram ботом

### Топики для бота

Рекомендуемая структура:

```bash
# События регистрации пользователей
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic bot.user.registered \
  --partitions 1 \
  --replication-factor 1

# Команды от пользователей
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic bot.user.commands \
  --partitions 3 \
  --replication-factor 1

# Системные события
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic bot.system.events \
  --partitions 1 \
  --replication-factor 1
```

### Пример Producer в боте

```go
// Отправка события регистрации
producer := kafka.NewProducer([]string{"localhost:9092"}, "bot.user.registered")

event := UserRegisteredEvent{
    UserID:    user.ID,
    TgID:      user.TgID,
    Username:  user.Username,
    Timestamp: time.Now(),
}

data, _ := json.Marshal(event)
producer.SendMessage(ctx, user.ID.String(), string(data))
```

### Пример Consumer

```go
// Обработка событий регистрации
consumer := kafka.NewConsumer(
    []string{"localhost:9092"},
    "bot.user.registered",
    "analytics-service",
)

for {
    msg, _ := consumer.ReadMessage(ctx)
    var event UserRegisteredEvent
    json.Unmarshal(msg.Value, &event)
    
    // Обработка события
    log.Printf("New user: %s (ID: %s)", event.Username, event.UserID)
}
```

## Мониторинг

### Логи

```bash
# Следить за логами в реальном времени
docker logs -f kafka

# Последние 100 строк
docker logs --tail 100 kafka

# С временными метками
docker logs -t kafka
```

### Метрики

```bash
# Статус брокера
docker exec -it kafka kafka-broker-api-versions.sh \
  --bootstrap-server localhost:9092

# Производительность топика
docker exec -it kafka kafka-run-class.sh kafka.tools.GetOffsetShell \
  --broker-list localhost:9092 \
  --topic test-topic
```

## Конфигурация Docker Compose

```yaml
services:
  kafka:
    image: bitnamilegacy/kafka:4.0.0-debian-12-r10
    container_name: kafka
    ports:
      - "9092:9092"    # Клиентский порт
      - "9093:9093"    # Контроллер
    environment:
      # KRaft режим
      - KAFKA_CFG_PROCESS_ROLES=broker,controller
      - KAFKA_CFG_NODE_ID=1
      
      # Координация
      - KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=1@localhost:9093
      
      # Listeners
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
      
      # Security
      - KAFKA_CFG_INTER_BROKER_LISTENER_NAME=PLAINTEXT
      - KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
      - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT
    volumes:
      - ./volumes:/bitnami/kafka
```

## Troubleshooting

### Kafka не запускается

```bash
# Проверить логи
docker logs kafka

# Проверить порты
netstat -an | grep -E '9092|9093'

# Пересоздать контейнер
docker-compose down
docker-compose up -d
```

### Ошибка "Connection refused"

Убедитесь, что:
1. Kafka запущен: `docker ps | grep kafka`
2. Порт доступен: `telnet localhost 9092`
3. В коде используется правильный адрес: `localhost:9092`

### Топик не создается

```bash
# Проверить права доступа к volumes
ls -la volumes/

# Очистить данные и пересоздать
docker-compose down -v
rm -rf volumes/
docker-compose up -d
```

### Высокое потребление памяти

Добавить лимиты в `docker-compose.yml`:

```yaml
services:
  kafka:
    # ...
    deploy:
      resources:
        limits:
          memory: 1G
        reservations:
          memory: 512M
```

## Production советы

### Конфигурация для продакшена

```yaml
environment:
  # Увеличить retention
  - KAFKA_CFG_LOG_RETENTION_HOURS=168  # 7 дней
  
  # Сжатие
  - KAFKA_CFG_COMPRESSION_TYPE=snappy
  
  # Производительность
  - KAFKA_CFG_NUM_NETWORK_THREADS=8
  - KAFKA_CFG_NUM_IO_THREADS=8
  
  # Репликация (для кластера)
  - KAFKA_CFG_DEFAULT_REPLICATION_FACTOR=3
  - KAFKA_CFG_MIN_INSYNC_REPLICAS=2
```

### Backup

```bash
# Бэкап данных
tar -czf kafka-backup-$(date +%Y%m%d).tar.gz volumes/

# Восстановление
docker-compose down
tar -xzf kafka-backup-20251129.tar.gz
docker-compose up -d
```

### Мониторинг (Prometheus + Grafana)

Добавить JMX exporter для метрик Kafka и использовать Grafana дашборды.

## Полезные ссылки

- [Kafka Documentation](https://kafka.apache.org/documentation/)
- [Bitnami Kafka Docker Image](https://github.com/bitnami/containers/tree/main/bitnami/kafka)
- [KRaft Mode Guide](https://kafka.apache.org/documentation/#kraft)

## Следующие шаги

1. ✅ Запустить Kafka
2. ✅ Создать топики для бота
3. 🔄 Интегрировать Producer в Telegram бота
4. 🔄 Добавить Consumer для аналитики
5. 📊 Настроить мониторинг