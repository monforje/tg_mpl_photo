package env

import (
	"os"
	"tgbot/pkg/errorx"
)

type Env struct {
	TgToken     string
	PostgresDSN string
}

func New() (*Env, error) {
	// godotenv.Load()

	t := os.Getenv("TG_TOKEN")
	if t == "" {
		return nil, errorx.ErrTgTokenEmpty
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		return nil, errorx.ErrPostgresDSNEmpty
	}

	return &Env{
		TgToken:     t,
		PostgresDSN: dsn,
	}, nil
}
