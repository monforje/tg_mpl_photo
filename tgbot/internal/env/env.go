package env

import (
	"fmt"
	"os"
	"tgbot/pkg/errorx"
	"tgbot/pkg/logx"

	"github.com/joho/godotenv"
)

type Env struct {
	TgToken     string
	PostgresDSN string
}

func New() (*Env, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	t := os.Getenv("TG_TOKEN")
	if t == "" {
		return nil, errorx.ErrTgTokenEmpty
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		return nil, errorx.ErrPostgresDSNEmpty
	}

	logx.Info("env loaded successfully")

	return &Env{
		TgToken:     t,
		PostgresDSN: dsn,
	}, nil
}
