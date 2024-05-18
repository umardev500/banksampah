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

type WasteDepositCreateSeed struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	WasteTypeID uuid.UUID `json:"waste_type_id"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
}

func (s *Seeder) WasteDepositSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding waste deposits...")
	filePath := "database/seeder/data/waste_deposit_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read waste deposits data file")
		return err
	}

	var rows []WasteDepositCreateSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal waste deposits data")
		return err
	}

	copyCount, err := q.CopyFrom(
		ctx,
		pgx.Identifier{"waste_deposits"},
		[]string{"id", "user_id", "waste_type_id", "quantity", "description"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]interface{}, error) {
			return []interface{}{
				rows[i].ID,
				rows[i].UserID,
				rows[i].WasteTypeID,
				rows[i].Quantity,
				rows[i].Description,
			}, nil
		}),
	)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to seed waste deposits")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding waste deposits... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) WasteDepositDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping waste deposits...")
	filePath := "database/seeder/data/waste_deposit_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read waste deposits data file")
		return err
	}

	var rows []WasteDepositCreateSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal waste deposits data")
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM waste_deposits WHERE id = ANY($1)"
	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to drop waste deposits")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping waste deposits... ✅")

	return nil
}
