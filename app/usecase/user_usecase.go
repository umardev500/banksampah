package usecase

import "github.com/umardev500/banksampah/domain"

type userUc struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUc{
		repo: repo,
	}
}
