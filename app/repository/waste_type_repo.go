package repository

import (
	"context"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type wasteTypeRepo struct {
	pgxConfig *config.PgxConfig
}

func NewWasteTypeRepo(pgxConfig *config.PgxConfig) domain.WasteTypeRepository {
	return &wasteTypeRepo{
		pgxConfig: pgxConfig,
	}
}

func (repo *wasteTypeRepo) Find(ctx context.Context) ([]model.WasteType, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM waste_types
	`
	rows, err := queries.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	var wasteTypes []model.WasteType
	for rows.Next() {
		wasteType := model.WasteType{}
		err := rows.Scan(
			&wasteType.ID,
			&wasteType.Name,
			&wasteType.Point,
			&wasteType.Description,
			&wasteType.CreatedAt,
			&wasteType.UpdatedAt,
			&wasteType.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		wasteTypes = append(wasteTypes, wasteType)
	}
	return wasteTypes, nil
}
