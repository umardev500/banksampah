package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/config"
)

type Migration struct {
	M *migrate.Migrate
}

func NewMigrate() *Migration {
	conn := config.NewPostgress()
	driver, err := postgres.WithInstance(conn.DBRaw, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("error driver connecting to postgres: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migration/files",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal().Msgf("error creating new migrate instance: %v", err)
	}

	log.Info().Msgf("🎉 Migration ready!")

	return &Migration{
		M: m,
	}
}

func (m *Migration) Up() {
	log.Info().Msgf("📦 Migrating up...")
	if err := m.M.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msgf("error migrating up: %v", err)
	}
	log.Info().Msgf("📦 Migrated up successfully!")
}

func (m *Migration) Down() {
	log.Info().Msgf("📦 Migrating down...")
	if err := m.M.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msgf("error migrating down: %v", err)
	}
	log.Info().Msgf("📦 Migrated down successfully!")
}
