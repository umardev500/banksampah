package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteDepoRepository struct {
	pgxConfig *config.PgxConfig
}

func NewWasteDepoRepository(pgxConfig *config.PgxConfig) domain.WasteDepoRepository {
	return &wasteDepoRepository{
		pgxConfig: pgxConfig,
	}
}

func (repo *wasteDepoRepository) DeleteByID(ctx context.Context, id string) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		DELETE FROM waste_deposits WHERE id=$1
	`
	result, err := queries.Exec(ctx, sql, id)
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}

func (repo *wasteDepoRepository) FindByID(ctx context.Context, id string) (wd *model.WasteDepo, err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM waste_deposits WHERE id = $1
	`
	var depo model.WasteDepo
	err = queries.QueryRow(ctx, sql, id).Scan(
		&depo.ID,
		&depo.UserID,
		&depo.WalletID,
		&depo.WasteTypeID,
		&depo.Quantity,
		&depo.Description,
		&depo.Status,
		&depo.CreatedAt,
		&depo.UpdatedAt,
		&depo.DeletedAt,
		&depo.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	wd = &depo

	return
}

func (repo *wasteDepoRepository) ConfirmDeposit(ctx context.Context, payload model.WasteDepoConfirmRequest) (*model.WasteDepo, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		UPDATE waste_deposits SET
	`
	rawSql, args := util.BuildUpdateQuery(sql, payload, []types.Filter{
		{
			Field:    "id",
			Operator: "=",
			Value:    payload.ID,
		},
		{
			Field:           "status",
			Operator:        "=",
			Value:           string(model.WasteDepoStatusUnConfirmed),
			LogicalOperator: "AND",
		},
	})
	rawSql = rawSql + " " + "RETURNING *"

	var depo model.WasteDepo
	err := queries.QueryRow(ctx, rawSql, args...).Scan(
		&depo.ID,
		&depo.UserID,
		&depo.WalletID,
		&depo.WasteTypeID,
		&depo.Quantity,
		&depo.Description,
		&depo.Status,
		&depo.CreatedAt,
		&depo.UpdatedAt,
		&depo.DeletedAt,
		&depo.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &depo, nil
}

func (repo *wasteDepoRepository) Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) (err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO waste_deposits (id, user_id, wallet_id, waste_type_id, quantity, "description", created_by) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = queries.Exec(
		ctx,
		sql,
		payload.ID,
		payload.UserID,
		payload.WalletID,
		payload.WasteTypeID,
		payload.Quantity,
		payload.Description,
		payload.CreatedBy,
	)

	return
}
