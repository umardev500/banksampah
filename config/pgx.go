package config

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/constant"
	"github.com/umardev500/banksampah/util"
)

type PgxQuery interface {
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type PgxConfig struct {
	Pool  *pgxpool.Pool
	Query PgxQuery
}

var (
	pgInstance *PgxConfig
	once       sync.Once
)

func NewPgx() *PgxConfig {
	once.Do(func() {
		logger := util.NewLogger()
		log.Info().Msg("🐘 Connecting to Postgres...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conf := &PgxConfig{}

		host := os.Getenv("POSTGRES_HOST")
		port := os.Getenv("POSTGRES_PORT")
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DATABASE")
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		dbpool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			log.Fatal().Msgf("error connecting to postgres: %v", err)
		}

		logger.UplineClearPrev()
		log.Info().Msg("🤝 Postgres ping....")

		if err := dbpool.Ping(ctx); err != nil {
			log.Fatal().Msgf("error pinging postgres: %v", err)
		}

		logger.UplineClearPrev()
		log.Info().Msg("👋 Postgres connected successfully!")
		conf.Pool = dbpool

		pgInstance = conf
	})

	return pgInstance
}

func (conf *PgxConfig) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := conf.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // Rollback transaction if not committed

	ctx = context.WithValue(ctx, constant.Tx, tx)
	err = fn(ctx)
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// Get connection type from context
func (conf *PgxConfig) TrOrDB(ctx context.Context) PgxQuery {
	if tx, ok := ctx.Value(constant.Tx).(PgxQuery); ok {
		return tx
	}

	return conf.Pool
}
