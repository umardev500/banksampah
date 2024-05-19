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

func (repo *walletRepo) Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO wallets (id, "name", "description") VALUES ($1, $2, $3);
	`
	_, err := queries.Exec(ctx, sql, payload.ID, payload.Name, payload.Description)
	return err
}
