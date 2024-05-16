package seeder

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/util"
)

type CreateWasteTypeSeed struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Point       float64   `json:"point"`
	Description string    `json:"description"`
}

func (s *Seeder) WasteTypeSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding waste types...")
	filePath := "database/seeder/data/waste_types_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read waste types data file")
		return err
	}

	var rows []CreateWasteTypeSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal waste types data")
		return err
	}

	copyCount, err := q.CopyFrom(
		ctx,
		pgx.Identifier{"waste_types"},
		[]string{"id", "name", "point", "description"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			return []any{
				rows[i].ID,
				rows[i].Name,
				rows[i].Point,
				rows[i].Description,
			}, nil
		}),
	)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to seed waste types")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding waste types... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) WasteTypeDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping waste types...")
	filePath := "database/seeder/data/waste_types_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read waste types data file")
		return err
	}

	var rows []CreateWasteTypeSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal waste types data")
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM waste_types WHERE id = ANY($1)"
	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to drop waste types")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping waste types... ✅")

	return nil
}
