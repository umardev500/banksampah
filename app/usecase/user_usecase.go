package usecase

import (
	"context"

	"github.com/gofiber/fiber/v3"
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

	err := uc.repo.Create(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			return response
		}

		return util.MakeResponse(
			payload.ID,
			fiber.StatusInternalServerError,
			fiber.ErrInternalServerError.Message,
			nil,
		)
	}

	return util.MakeResponse(payload.ID, 200, "Create user successfuly", nil)
}
