package repository

import (
	"context"

	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type userRepo struct {
	pgxConfig *config.PgxConfig
}

func NewUserRepo(pgxConfig *config.PgxConfig) domain.UserRepository {
	return &userRepo{
		pgxConfig: pgxConfig,
	}
}

func (u *userRepo) Create(ctx context.Context, payload model.CreateUser) error {
	queries := u.pgxConfig.TrOrDB(ctx)
	sql := `--sql
		INSERT INTO users (id, email, username, "password") values ($1, $2, $3, $4)
	`
	_, err := queries.Exec(ctx, sql, payload.ID, payload.Email, payload.Username, payload.Password)
	return err
}
