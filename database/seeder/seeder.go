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
		// Feature
		err = s.FeatureDown(ctx)
		if err != nil {
			return err
		}

		err = s.FeatureSeeds(ctx)
		if err != nil {
			return err
		}

		// Role
		err = s.RoleDown(ctx)
		if err != nil {
			return err
		}

		err = s.RoleSeeds(ctx)
		if err != nil {
			return err
		}

		// User
		err = s.UserDown(ctx)
		if err != nil {
			return err
		}

		err = s.UserSeeds(ctx)
		if err != nil {
			return err
		}

		// User Role
		err = s.UserRoleDown(ctx)
		if err != nil {
			return err
		}

		err = s.UserRoleSeeds(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}
