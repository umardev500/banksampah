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

	return nil
}
