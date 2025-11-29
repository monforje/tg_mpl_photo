# Telegram Bot

Telegram бот на Go с чистой архитектурой, регистрацией пользователей и интеграцией с PostgreSQL.

## Возможности

- ✅ Регистрация пользователей через `/start`
- ✅ Хранение данных в PostgreSQL
- ✅ Миграции базы данных (Goose)
- ✅ Graceful shutdown
- ✅ Структурированное логирование (slog)
- ✅ Clean Architecture (handlers → services → repos)
- 🔄 Готов к интеграции с Kafka

## Архитектура

```
cmd/
├── tgbot/          # Точка входа бота
└── goose/          # Утилита для миграций

internal/
├── app/            # Инициализация приложения
├── core/           # Бизнес-логика
│   ├── model/      # Модели данных
│   └── repo/       # Интерфейсы репозиториев
├── env/            # Загрузка окружения
├── postgres/       # Работа с БД
│   └── repoimpl/   # Реализация репозиториев
├── service/        # Бизнес-сервисы
└── tg/             # Telegram слой
    ├── bot/        # Инициализация бота
    ├── handler/    # Обработчики команд
    └── router/     # Роутинг команд

pkg/                # Переиспользуемые пакеты
├── errorx/         # Кастомные ошибки
├── goose/          # Обертка над goose
├── logx/           # Логирование
└── tg/             # Telegram утилиты
    ├── message/    # Текстовые сообщения
    └── sticker/    # Стикеры

migrations/         # SQL миграции
```

## Быстрый старт

### 1. Настройка окружения

Создайте `.env` файл:

```bash
cp .env.example .env
```

Заполните переменные:

```env
TG_TOKEN=123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11
POSTGRES_DSN=postgres://user:password@localhost:5432/botdb?sslmode=disable
```

**Получение токена:**
1. Напишите [@BotFather](https://t.me/BotFather)
2. Отправьте `/newbot`
3. Следуйте инструкциям
4. Скопируйте токен в `.env`

### 2. Установка зависимостей

```bash
go mod download
```

### 3. Запуск PostgreSQL

С помощью Docker:

```bash
docker run -d \
  --name postgres-bot \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=botdb \
  -p 5432:5432 \
  postgres:16-alpine
```

### 4. Применение миграций

```bash
go run cmd/goose/main.go
```

### 5. Запуск бота

```bash
go run cmd/tgbot/main.go
```

Лог успешного запуска:

```
time=... level=INFO msg="env loaded successfully"
time=... level=INFO msg="connecting to postgres"
time=... level=INFO msg="postgres connection established"
time=... level=INFO msg="bot initialized successfully"
time=... level=INFO msg="app initialized successfully"
time=... level=INFO msg="app is starting"
time=... level=INFO msg="bot is starting"
```

## Использование

### Регистрация пользователя

1. Найдите бота в Telegram
2. Отправьте `/start`
3. Получите подтверждение регистрации

**Ответ при первой регистрации:**
```
[Стикер успеха]
You have successfully registered!
```

**При повторной попытке:**
```
[Стикер успеха]
You are already registered.
```

## База данных

### Схема `users`

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    tg_id BIGINT NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_tg_id ON users(tg_id);
```

### Работа с миграциями

```bash
# Применить все миграции
go run cmd/goose/main.go

# Откатить последнюю миграцию
goose -dir migrations postgres "$POSTGRES_DSN" down

# Статус миграций
goose -dir migrations postgres "$POSTGRES_DSN" status

# Создать новую миграцию
goose -dir migrations create add_user_balance sql
```

## Разработка

### Добавление новой команды

1. **Создать handler** в `internal/tg/handler/`:

```go
package handler

type MyHandler struct {
    service *service.MyService
}

func (h *MyHandler) HandleMyCommand(c tele.Context) error {
    // Ваша логика
    return c.Send("Response")
}
```

2. **Зарегистрировать в router** (`internal/tg/router/router.go`):

```go
func (r *Router) Setup() {
    r.bot.Handle("/mycommand", r.myHandler.HandleMyCommand)
}
```

3. **Добавить service** (если нужна бизнес-логика):

```go
// internal/service/my_service.go
type MyService struct {
    repo repo.MyRepo
}

func (s *MyService) DoSomething() error {
    // Бизнес-логика
}
```

### Добавление нового репозитория

1. **Интерфейс** в `internal/core/repo/`:

```go
type MyRepo interface {
    GetSomething(id uuid.UUID) (*model.Something, error)
}
```

2. **Реализация** в `internal/postgres/repoimpl/`:

```go
type MyRepoImpl struct {
    pool *pgxpool.Pool
}

func (r *MyRepoImpl) GetSomething(id uuid.UUID) (*model.Something, error) {
    // SQL запрос
}
```

## Конфигурация

### Переменные окружения

| Переменная | Описание | Пример |
|------------|----------|--------|
| `TG_TOKEN` | Токен Telegram бота | `123456:ABC-DEF...` |
| `POSTGRES_DSN` | Строка подключения к PostgreSQL | `postgres://user:pass@localhost/db` |

### Логирование

Формат: Text (можно переключить на JSON в `pkg/logx/logx.go`)

Уровни:
- `INFO` - общая информация
- `ERROR` - ошибки с контекстом
- `FATAL` - критические ошибки (с выходом)

Пример лога регистрации:

```
level=INFO msg="user registered successfully" username=john_doe tg_id=123456789
```

## Интеграция с Kafka (планируется)

### Producer для событий

```go
// Отправка события регистрации в Kafka
func (r *RegService) Reg(tgID int64, username string) error {
    // ... создание пользователя ...
    
    // Отправить событие в Kafka
    event := Event{
        Type: "user.registered",
        UserID: id,
        TgID: tgID,
        Timestamp: time.Now(),
    }
    r.kafkaProducer.SendEvent(ctx, event)
}
```

### Consumer для обработки команд

```go
// Обработка команд из Kafka очереди
func (b *Bot) ProcessKafkaCommands(ctx context.Context) {
    for {
        msg := consumer.ReadMessage(ctx)
        // Выполнить команду через бота
    }
}
```

## Тестирование

```bash
# Запустить тесты
go test ./...

# С покрытием
go test -cover ./...

# Конкретный пакет
go test ./internal/service/...
```

## Production

### Docker build

```bash
# Билд бота
docker build -t tgbot:latest -f Dockerfile .

# Запуск
docker run -d \
  --name tgbot \
  --env-file .env \
  tgbot:latest
```

### Рекомендации

- Используйте переменные окружения для секретов
- Настройте мониторинг и алертинг
- Добавьте health checks
- Настройте ротацию логов
- Используйте connection pooling для PostgreSQL
- Добавьте rate limiting для команд

## Зависимости

```go
require (
    github.com/joho/godotenv v1.5.1           // Загрузка .env
    gopkg.in/telebot.v4 v4.0.0-beta.7         // Telegram Bot API
    github.com/google/uuid v1.6.0             // UUID генерация
    github.com/jackc/pgx/v5 v5.7.6            // PostgreSQL драйвер
    github.com/pressly/goose/v3 v3.26.0       // Миграции
)
```

## Troubleshooting

### Бот не отвечает

```bash
# Проверить токен
echo $TG_TOKEN

# Проверить подключение к Telegram API
curl https://api.telegram.org/bot$TG_TOKEN/getMe
```

### Ошибка подключения к PostgreSQL

```bash
# Проверить доступность
psql "$POSTGRES_DSN"

# Проверить порт
netstat -an | grep 5432

# Логи PostgreSQL
docker logs postgres-bot
```

### Миграции не применяются

```bash
# Проверить статус
goose -dir migrations postgres "$POSTGRES_DSN" status

# Принудительно применить
goose -dir migrations postgres "$POSTGRES_DSN" up
```

## Roadmap

- [ ] Добавить команды для работы с профилем
- [ ] Интегрировать Kafka Producer/Consumer
- [ ] Добавить middleware для логирования
- [ ] Реализовать rate limiting
- [ ] Добавить unit тесты
- [ ] Настроить CI/CD
- [ ] Добавить метрики (Prometheus)
- [ ] Реализовать админ-панель

## Контакты

При возникновении вопросов создайте Issue в репозитории.