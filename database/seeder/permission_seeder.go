package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/util"
)

type CreatePermSeed struct {
	ID          uuid.UUID `json:"id"`
	FeatureID   uuid.UUID `json:"feature_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func (s *Seeder) PermissionSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding permissions...")
	filePaths := []string{
		"database/seeder/data/perm_feature.json",
		"database/seeder/data/perm_permission.json",
		"database/seeder/data/perm_role_man.json",
		"database/seeder/data/perm_user_man.json",
	}

	var allContent []CreatePermSeed

	for _, filePath := range filePaths {
		var each []CreatePermSeed
		if err := util.ParseJSONFile(filePath, &each); err != nil {
			log.Error().Msgf("error parsing json file: %v", err)

			return err
		}

		allContent = append(allContent, each...)
	}

	copyCount, err := q.CopyFrom(
		ctx,
		pgx.Identifier{"permissions"},
		[]string{"id", "feature_id", "name", "description", "status"},
		pgx.CopyFromSlice(len(allContent), func(i int) ([]any, error) {
			values := []any{
				allContent[i].ID,
				allContent[i].FeatureID,
				allContent[i].Name,
				allContent[i].Description,
				allContent[i].Status,
			}

			return values, nil
		}),
	)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Seeding permissions... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding permissions... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) PermissionDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping permissions...")

	filePaths := []string{
		"database/seeder/data/perm_feature.json",
		"database/seeder/data/perm_permission.json",
		"database/seeder/data/perm_role_man.json",
		"database/seeder/data/perm_user_man.json",
	}

	var allContent []CreatePermSeed

	for _, filePath := range filePaths {
		var each []CreatePermSeed
		if err := util.ParseJSONFile(filePath, &each); err != nil {
			log.Error().Msgf("error parsing json file: %v", err)

			return err
		}

		allContent = append(allContent, each...)
	}

	var idsToDelete []uuid.UUID
	for _, val := range allContent {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM permissions WHERE id = ANY($1)"

	_, err := q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Dropping permissions... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(150 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping permissions... ✅")

	return nil
}
