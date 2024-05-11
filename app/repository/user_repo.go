package repository

import (
	"context"
	"errors"

	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type userRepo struct{}

func NewUserRepo() domain.UserRepository {
	return &userRepo{}
}

func (u *userRepo) Create(ctx context.Context, payload model.CreateUser) error {
	return errors.New("h")
}
