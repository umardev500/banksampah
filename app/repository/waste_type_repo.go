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

type wasteTypeRepo struct {
	pgxConfig *config.PgxConfig
}

func NewWasteTypeRepo(pgxConfig *config.PgxConfig) domain.WasteTypeRepository {
	return &wasteTypeRepo{
		pgxConfig: pgxConfig,
	}
}

func (repo *wasteTypeRepo) FindByID(ctx context.Context, id string) (wt *model.WasteType, err error) {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		SELECT * FROM waste_types WHERE id = $1
	`

	var result model.WasteType
	err = queries.QueryRow(ctx, sql, id).Scan(
		&result.ID,
		&result.Name,
		&result.Point,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	)

	wt = &result

	return
}

func (repo *wasteTypeRepo) CreateWithVersion(ctx context.Context, payload model.WasteTypeCreateWithVersion) (*model.WasteType, error) {
	var err error
	var wt *model.WasteType
	err = repo.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		err = repo.createVersion(ctx, payload)
		if err != nil {
			return err
		}
		wt, err = repo.Create(ctx, payload)

		return err
	})

	return wt, nil
}

func (repo *wasteTypeRepo) createVersion(ctx context.Context, payload model.WasteTypeCreateWithVersion) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO wt_versions (id, wt_id, "name", "point", "description", created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := queries.Exec(
		ctx,
		sql,
		payload.VERSIONID,
		payload.SOURCEID,
		payload.Name,
		payload.Point,
		payload.Description,
		payload.CreatedBy,
	)
	return err
}

func (repo *wasteTypeRepo) Create(ctx context.Context, payload model.WasteTypeCreateWithVersion) (*model.WasteType, error) {
	var result model.WasteType

	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO waste_types (id, "name", "point", "description", "version_id", created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *
	`

	// _, err := queries.Exec(ctx, sql, payload.ID, payload.Name, payload.Point, payload.Description)
	row := queries.QueryRow(
		ctx,
		sql,
		payload.SOURCEID,
		payload.Name,
		payload.Point,
		payload.Description,
		payload.VERSIONID,
		payload.CreatedBy,
	)

	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Point,
		&result.Description,
		&result.VersionID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
		&result.DeletedBy,
		&result.UpdatedBy,
		&result.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (repo *wasteTypeRepo) UpdateByIDWithVersion(ctx context.Context, payload model.WasteTypeUpdateWithVersionRequest) error {
	var err error
	err = repo.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		err = repo.createVersion(ctx, model.WasteTypeCreateWithVersion{
			SOURCEID:    payload.SOURCEID,
			VERSIONID:   payload.VERSIONID,
			Name:        payload.Name,
			Point:       payload.Point,
			Description: payload.Description,
			CreatedBy:   payload.UpdatedBy,
		})
		if err != nil {
			return err
		}
		err = repo.UpdateByID(ctx, model.WasteTypeUpdateWithVersionRequest{
			VERSIONID:   payload.VERSIONID,
			Name:        payload.Name,
			Point:       payload.Point,
			Description: payload.Description,
			UpdatedBy:   payload.UpdatedBy,
		})

		return nil
	})

	return err
}

func (repo *wasteTypeRepo) UpdateByID(ctx context.Context, payload model.WasteTypeUpdateWithVersionRequest) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		UPDATE waste_types SET
	`
	rawQuery, args := util.BuildUpdateQuery(sql, payload, []types.Filter{
		{
			Field:    "id",
			Operator: "=",
			Value:    payload.SOURCEID,
		},
	})
	if args == nil {
		return nil
	}

	_, err := queries.Exec(ctx, rawQuery, args...)
	return err
}

func (repo *wasteTypeRepo) DeleteByID(ctx context.Context, id string) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		DELETE FROM waste_types WHERE id=$1
	`
	result, err := queries.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
}

func (repo *wasteTypeRepo) SoftDeleteByID(ctx context.Context, deletedBy, id string) error {
	queries := repo.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		UPDATE waste_types SET deleted_at = now(), deleted_by = $1 WHERE id = $2
	`
	result, err := queries.Exec(ctx, sql, deletedBy, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return err
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
			&wasteType.VersionID,
			&wasteType.CreatedAt,
			&wasteType.UpdatedAt,
			&wasteType.DeletedAt,
			&wasteType.DeletedBy,
			&wasteType.UpdatedBy,
			&wasteType.CreatedBy,
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
