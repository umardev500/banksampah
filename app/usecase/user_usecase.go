package usecase

import (
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type userUc struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUc{
		repo: repo,
	}
}

func (uc *userUc) Create(payload model.CreateUser) error {
	return nil
}
