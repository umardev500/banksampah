package seeder

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/util"
)

type CreateFeatureSeed struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func (s *Seeder) FeatureSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding features...")
	filePath := "database/seeder/data/feature_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateFeatureSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	copyCount, err := q.CopyFrom(
		context.Background(),
		pgx.Identifier{"features"},
		[]string{"id", "name", "description", "status"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			values := []any{
				rows[i].ID,
				rows[i].Name,
				rows[i].Description,
				rows[i].Status,
			}

			return values, nil
		}),
	)

	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Seeding features... 🚧")
		fmt.Println(err)
		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding features... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) FeatureDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping features...")
	filePath := "database/seeder/data/feature_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateFeatureSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM features WHERE id = ANY($1)"

	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Dropping features... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping features... ✅")

	return nil
}
