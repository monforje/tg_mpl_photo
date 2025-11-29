package main

import (
	"tgbot/internal/env"
	"tgbot/pkg/goose"
	"tgbot/pkg/logx"
)

func main() {
	e, err := env.New()
	if err != nil {
		logx.Fatal("failed to load environment", "error", err)
	}

	if err := goose.Migrate(e.PostgresDSN); err != nil {
		logx.Fatal("migration failed", "error", err)
	}
}
