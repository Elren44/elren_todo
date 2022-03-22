package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Elren44/elog"
	"github.com/Elren44/elren_todo/internal/config"
	"github.com/Elren44/elren_todo/pkg/utils"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	logger := elog.InitLogger(elog.ConsoleOutput)
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, connStr)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if err != nil {
		logger.Fatalf("error do with tries postrgesql: %v", err)
	}
	return pool, err
}
