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
		INSERT INTO waste_deposits (id, user_id, waste_type_id, quantity, "description", created_by) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = queries.Exec(
		ctx,
		sql,
		payload.ID,
		payload.UserID,
		payload.WasteTypeID,
		payload.Quantity,
		payload.Description,
		payload.CreatedBy,
	)

	return
}
