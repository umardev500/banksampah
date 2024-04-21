package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type PostgresConfig struct {
	Conn *sqlx.DB
}

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

	conf.Conn = db

	// ping log with emoji
	log.Info().Msg("🤝 Postgres ping....")

	if err := db.Ping(); err != nil {
		log.Fatal().Msgf("error pinging postgres: %v", err)
	}

	log.Info().Msg("👋 Postgres connected successfully!")

	return conf
}
