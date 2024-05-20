package repository

import (
	"context"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type wasteDepoRepository struct {
	pgxConfig *config.PgxConfig
}

func NewWasteDepoRepository(pgxConfig *config.PgxConfig) domain.WasteDepoRepository {
	return &wasteDepoRepository{
		pgxConfig: pgxConfig,
	}
}

func (repo *wasteDepoRepository) Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) (err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO waste_deposits (id, user_id, quantity, "description") VALUES ($1,$2,$3,$4)
	`

	_, err = queries.Exec(ctx, sql, payload.ID, payload.UserID, payload.Quantity, payload.Description)

	return
}
