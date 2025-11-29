package goose

import (
	"database/sql"
	"tgbot/pkg/logx"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Migrate(dsn string) error {
	logx.Info("starting database migrations")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logx.Fatal("failed to connect to database", "error", err)
		return err
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		logx.Fatal("failed to set goose dialect", "error", err)
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		logx.Fatal("failed to run migrations", "error", err)
		return err
	}

	logx.Info("migrations completed successfully")
	return nil
}
