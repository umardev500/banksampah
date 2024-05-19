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

type WalletCreateSeed struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
}

func (s *Seeder) WalletSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding wallets...")
	filePath := "database/seeder/data/wallet_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read wallets data file")
		return err
	}

	var rows []WalletCreateSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal wallets data")
		return err
	}

	copyCount, err := q.CopyFrom(
		ctx,
		pgx.Identifier{"wallets"},
		[]string{"id", "user_id", "name", "balance", "description", "type"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]interface{}, error) {
			return []interface{}{
				rows[i].ID,
				rows[i].UserID,
				rows[i].Name,
				rows[i].Balance,
				rows[i].Description,
				rows[i].Type,
			}, nil
		}),
	)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to seed wallets")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding wallets... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) WalletDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping wallets...")
	filePath := "database/seeder/data/wallet_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read wallets data file")
		return err
	}

	var rows []WalletCreateSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal wallets data")
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM wallets WHERE id = ANY($1)"
	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Err(err).Msg("Failed to drop wallets")
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping wallets... ✅")

	return nil
}
