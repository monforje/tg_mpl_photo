package psql

import (
	"context"
	"tgbot/pkg/logx"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Psql struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) (*Psql, error) {
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

	return &Psql{
		Pool: pool,
	}, nil
}

func (p *Psql) Stop() {
	p.Pool.Close()
	logx.Info("postgres connection closed")
}
