package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/constant"
)

type Query interface {
	// Methods from sql.Tx
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row

	// Methods from sqlx.Tx
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	DriverName() string
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type PostgresConfig struct {
	DB    *sqlx.DB
	DBRaw *sql.DB
	Query Query
}

type TxnFn func(ctx context.Context) error

func NewPostgress() *PostgresConfig {
	log.Info().Msg("🐘 Connecting to Postgres...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conf := &PostgresConfig{}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Fatal().Msgf("error connecting to postgres: %v", err)
	}

	conf.DB = db
	conf.DBRaw = db.DB

	// ping log with emoji
	log.Info().Msg("🤝 Postgres ping....")

	if err := db.Ping(); err != nil {
		log.Fatal().Msgf("error pinging postgres: %v", err)
	}

	log.Info().Msg("👋 Postgres connected successfully!")

	return conf
}

func (conf *PostgresConfig) WithTransaction(ctx context.Context, fn TxnFn) error {
	tx, err := conf.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, constant.Tx, tx)

	err = fn(ctx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (conf *PostgresConfig) TrOrDB(ctx context.Context) {
	if tx, ok := ctx.Value(constant.Tx).(Query); ok {
		fmt.Println("Is Query")
		conf.Query = tx
	}

	conf.Query = conf.DB
}
