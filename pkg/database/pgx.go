package database

import (
	"context"
	"log"

	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(cfg *config.Config) DB {
	log.Println(cfg.Credential.DB.URL)
	client, err := pgxpool.New(context.Background(), cfg.Credential.DB.URL)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}

	return &dbImpl{
		client: client,
	}
}

func (d *dbImpl) Begin(ctx context.Context) (pgx.Tx, error) {
	return d.client.Begin(ctx)
}

func (d *dbImpl) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return d.client.Exec(ctx, sql, arguments...)
}

func (d *dbImpl) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return d.client.Query(ctx, sql, args...)
}

func (d *dbImpl) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return d.client.QueryRow(ctx, sql, args...)
}
