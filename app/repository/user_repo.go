package repository

import "github.com/umardev500/banksampah/domain"

type userRepo struct{}

func NewUserRepo() domain.UserRepository {
	return &userRepo{}
}
