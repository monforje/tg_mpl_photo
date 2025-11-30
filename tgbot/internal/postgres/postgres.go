package postgres

import (
	"context"
	"fmt"
	"tgbot/pkg/logx"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping postgres pool: %w", err)
	}

	logx.Info("postgres connection established")

	return &Postgres{
		Pool: pool,
	}, nil
}

func (p *Postgres) Stop() {
	p.Pool.Close()
	logx.Info("postgres connection closed")
}
