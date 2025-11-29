package postgres

import (
	"context"
	"tgbot/pkg/logx"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Postgres, error) {
	logx.Info("connecting to postgres")

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
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
