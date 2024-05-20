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
		/* DOWN START */
		// Permission
		err = s.PermissionDown(ctx)
		if err != nil {
			return err
		}

		// feature
		err = s.FeatureDown(ctx)
		if err != nil {
			return err
		}
		// User Role
		err = s.UserRoleDown(ctx)
		if err != nil {
			return err
		}
		// Role
		err = s.RoleDown(ctx)
		if err != nil {
			return err
		}

		// Waste Deposit
		err = s.WasteDepositDown(ctx)
		if err != nil {
			return err
		}

		// Wallet
		err = s.WalletDown(ctx)
		if err != nil {
			return err
		}

		// User
		err = s.UserDown(ctx)
		if err != nil {
			return err
		}

		// Waste Type
		err = s.WasteTypeDown(ctx)
		if err != nil {
			return err
		}
		/* DOWN END */

		/* UP START */

		// Feature
		err = s.FeatureSeeds(ctx)
		if err != nil {
			return err
		}

		// Role
		err = s.RoleSeeds(ctx)
		if err != nil {
			return err
		}

		// User
		err = s.UserSeeds(ctx)
		if err != nil {
			return err
		}

		// User Role
		err = s.UserRoleSeeds(ctx)
		if err != nil {
			return err
		}

		// Permission
		err = s.PermissionSeeds(ctx)
		if err != nil {
			return err
		}

		// Waste Type
		err = s.WasteTypeSeeds(ctx)
		if err != nil {
			return err
		}

		// Wallet
		err = s.WalletSeeds(ctx)
		if err != nil {
			return err
		}

		// Waste Deposit
		err = s.WasteDepositSeeds(ctx)
		if err != nil {
			return err
		}

		/* END UP */

		return nil
	})
}
