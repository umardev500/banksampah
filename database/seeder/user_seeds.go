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

type CreateUserSeed struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email" validate:"required,email,min=7"`
	Username string    `json:"username" validate:"required,min=6"`
	Password string    `json:"password" validate:"required,min=8"`
}

func (s *Seeder) UserSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding users...")
	filePath := "database/seeder/data/user_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateUserSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	copyCount, err := q.CopyFrom(
		context.Background(),
		pgx.Identifier{"users"},
		[]string{"id", "email", "username", "password"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			values := []any{
				rows[i].ID,
				rows[i].Email,
				rows[i].Username,
				rows[i].Password,
			}

			return values, nil
		}),
	)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Seeding users... 🚧")
		fmt.Println(err)

		return err
	}
	time.Sleep(500 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeded %d users", copyCount)

	return nil
}

func (s *Seeder) UserDown(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Dropping users...")
	filePath := "database/seeder/data/user_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateUserSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	var idsToDelete []uuid.UUID
	for _, val := range rows {
		idsToDelete = append(idsToDelete, val.ID)
	}

	sql := "DELETE FROM users WHERE id = ANY($1)"

	_, err = q.Exec(ctx, sql, idsToDelete)
	if err != nil {
		logger.Upline()
		logger.FirstLine()
		log.Error().Msg("📦 Dropping users... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(500 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msg("📦 Users dropped successfully")

	return nil
}
