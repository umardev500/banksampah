package usecase

import "github.com/umardev500/banksampah/domain"

type userUc struct{}

func NewUserUsecase() domain.UserUsecase {
	return &userUc{}
}
