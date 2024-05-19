package repository

import (
	"context"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type walletRepo struct {
	pgxConfig *config.PgxConfig
}

func NewWalletRepository(pgxConfig *config.PgxConfig) domain.WalletRepository {
	return &walletRepo{
		pgxConfig: pgxConfig,
	}
}

func (repo *walletRepo) FindByID(ctx context.Context, id string) (model.Wallet, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM wallets WHERE id=$1
	`
	row := queries.QueryRow(ctx, sql, id)
	var result model.Wallet
	err := row.Scan(
		&result.ID,
		&result.UserD,
		&result.Name,
		&result.Balance,
		&result.Description,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)
	return result, err
}

func (repo *walletRepo) DeleteByID(ctx context.Context, id, userID string) error {
	sql := `--sql
		DELETE FROM wallets WHERE id=$1 AND "type" != 'master'
	`

	restoreBalanceToMaster := `--sql
		UPDATE wallets SET balance = balance + $1 WHERE "type" = 'master' AND user_id = $2
	`

	// Transaction
	err := repo.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		queries := repo.pgxConfig.TrOrDB(ctx)
		result, err := queries.Exec(ctx, sql, id)
		if err != nil {
			return err
		}

		rowsAffected := result.RowsAffected()

		// Update balance to master
		if rowsAffected > 0 {
			_, err = queries.Exec(ctx, restoreBalanceToMaster, rowsAffected, userID)
			if err != nil {
				return err
			}
		}

		return err
	})

	return err
}

func (repo *walletRepo) Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (model.Wallet, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO wallets (id, user_id, "name", "description", "type") VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	// result, err := queries.Exec(ctx, sql, payload.ID, payload.UserID, payload.Name, payload.Description)
	// fmt.Println(result)

	row := queries.QueryRow(ctx, sql, payload.ID, payload.UserID, payload.Name, payload.Description, payload.Type)
	var result model.Wallet
	err := row.Scan(
		&result.ID,
		&result.UserD,
		&result.Name,
		&result.Balance,
		&result.Description,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	return result, err
}
