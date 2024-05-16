package repository

import (
	"context"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteTypeRepo struct {
	pgxConfig *config.PgxConfig
}

func NewWasteTypeRepo(pgxConfig *config.PgxConfig) domain.WasteTypeRepository {
	return &wasteTypeRepo{
		pgxConfig: pgxConfig,
	}
}

func (repo *wasteTypeRepo) Find(ctx context.Context, params *types.QueryParam) (*model.FindWasteTypeResponse, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		WITH items_count AS (
			SELECT count(*) as total FROM waste_types
		)
		SELECT wt.*, ic.total FROM waste_types wt, items_count ic
	`

	queryRaw, args := util.BuildQuery(sql, params)
	queryRaw = util.StringTrimAnNoExtraSpace(util.RemoveSqlComment(queryRaw))

	rows, err := queries.Query(ctx, queryRaw, args...)
	if err != nil {
		return nil, err
	}
	var wasteTypes []model.WasteType
	var total int
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
			&total,
		)
		if err != nil {
			rows.Close()
			return nil, err
		}
		wasteTypes = append(wasteTypes, wasteType)
	}

	return &model.FindWasteTypeResponse{
		Total:      total,
		WasteTypes: wasteTypes,
	}, nil
}
