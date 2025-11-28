# Kafka Server

Kafka сервер в Docker контейнере с использованием KRaft (без Zookeeper).

## Запуск

```bash
docker-compose up -d
```

## Остановка

```bash
docker-compose down
```

## Конфигурация

- **Image:** `bitnamilegacy/kafka:4.0.0-debian-12-r10`
- **Порты:**
  - `9092` - PLAINTEXT (для клиентов)
  - `9093` - CONTROLLER (внутренний)
- **Режим:** KRaft (без Zookeeper)
- **Node ID:** 1

## Volumes

Данные Kafka сохраняются в `./volumes/`:
- `volumes/data/` - логи и партиции топиков
- `volumes/config/` - конфигурационные файлы

## Доступ

Kafka доступен на `localhost:9092`

## Команды для работы

### Создать топик

```bash
docker exec -it kafka kafka-topics.sh --create \
  --bootstrap-server localhost:9092 \
  --topic test-topic \
  --partitions 1 \
  --replication-factor 1
```

### Список топиков

```bash
docker exec -it kafka kafka-topics.sh --list \
  --bootstrap-server localhost:9092
```

### Удалить топик

```bash
docker exec -it kafka kafka-topics.sh --delete \
  --bootstrap-server localhost:9092 \
  --topic test-topic
```

## Логи

```bash
docker logs -f kafka
```