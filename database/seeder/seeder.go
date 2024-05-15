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
	s.Conn.WithTransaction(context.Background(), func(ctx context.Context) (err error) {
		err = s.RoleSeeds(ctx)
		if err != nil {
			return err
		}

		err = s.UserSeeds(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
