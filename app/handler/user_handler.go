package handler

import "github.com/umardev500/banksampah/domain"

type userH struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) domain.UserHandler {
	return &userH{
		uc: uc,
	}
}
