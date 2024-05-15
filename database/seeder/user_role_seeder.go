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

type CreateUserRoleSeed struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}

func (s *Seeder) UserRoleSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping user_roles...")
	filePath := "database/seeder/data/user_role_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateUserRoleSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	copyCount, err := q.CopyFrom(
		context.Background(),
		pgx.Identifier{"user_roles"},
		[]string{"user_id", "role_id"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			values := []any{
				rows[i].UserID,
				rows[i].RoleID,
			}

			return values, nil
		}),
	)

	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Seeding user_roles... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(500 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeding user_roles... ✅ (%d rows)", copyCount)

	return nil
}

func (s *Seeder) UserRoleDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping user_roles...")
	filePath := "database/seeder/data/user_role_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateUserRoleSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.UserID)
	}

	sql := "DELETE FROM user_roles WHERE user_id = ANY($1)"

	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Dropping user_roles... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(500 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Dropping user_roles... ✅")

	return nil
}
