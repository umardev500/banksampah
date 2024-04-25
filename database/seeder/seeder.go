package seeder

import (
	"context"

	"github.com/umardev500/banksampah/config"
)

type Seeder struct {
	Conn *config.PgxConfig
}

func NewSeeder() *Seeder {
	return &Seeder{
		Conn: config.NewPgx(),
	}
}

func (s *Seeder) Register() {
	s.Conn.WithTransaction(context.Background(), func(ctx context.Context) error {
		s.UserSeeds(ctx)

		return nil
	})
}
