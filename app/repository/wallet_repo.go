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

func (repo *walletRepo) DeleteByID(ctx context.Context, id string) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		DELETE FROM wallets WHERE id=$1
	`
	_, err := queries.Exec(ctx, sql, id)
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
		&result.Amount,
		&result.Description,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	return result, err
}
