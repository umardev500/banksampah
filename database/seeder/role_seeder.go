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

type CreateRoleSeed struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
}

func (s *Seeder) RoleSeeds(ctx context.Context) error {
	q := s.Conn.TrOrDB(ctx)
	logger := util.NewLogger()

	log.Info().Msg("📦 Seeding roles...")
	filePath := "database/seeder/data/role_data.json"
	f, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var rows []CreateRoleSeed
	err = json.Unmarshal(f, &rows)
	if err != nil {
		return err
	}

	copyCount, err := q.CopyFrom(
		context.Background(),
		pgx.Identifier{"roles"},
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
		log.Error().Msg("📦 Seeding roles... 🚧")
		fmt.Println(err)

		return err
	}

	time.Sleep(500 * time.Millisecond) // add delay

	logger.UplineClearPrev()

	log.Info().Msgf("📦 Seeded %d roles", copyCount)

	return nil
}
