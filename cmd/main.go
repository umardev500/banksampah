package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/app"
	"github.com/umardev500/banksampah/database/migration"
)

func init() {
	// log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := godotenv.Load(); err != nil {
		log.Fatal().Msgf("error loading .env file: %v", err)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()

	m := migration.NewMigrate()
	m.Up()

	v := validator.New()
	application := app.New(v)
	err := application.Run(ctx)
	if err != nil {
		log.Fatal().Msgf("error running app: %v", err)
	}

}
