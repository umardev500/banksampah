package repository

import (
	"context"
	"fmt"

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

func (repo *wasteTypeRepo) Find(ctx context.Context, params *types.QueryParam) ([]model.WasteType, error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM waste_types
	`

	queryRaw, args := util.BuildQuery(sql, params)
	queryRaw = util.StringTrimAnNoExtraSpace(util.RemoveSqlComment(queryRaw))

	fmt.Println(queryRaw)

	rows, err := queries.Query(ctx, queryRaw, args...)
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
			rows.Close()
			return nil, err
		}
		wasteTypes = append(wasteTypes, wasteType)
	}
	return wasteTypes, nil
}
