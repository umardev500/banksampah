package usecase

import (
	"context"

	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type userUc struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUc{
		repo: repo,
	}
}

func (uc *userUc) Create(ctx context.Context, payload model.CreateUser) util.Response {
	payload.ID = util.GenerateUUID()

	return util.MakeResponse(200, "Create user successfuly", nil)
}
