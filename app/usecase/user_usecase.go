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
	var err error

	payload.ID = util.GenerateUUID()
	payload.Password, err = util.GenerateBcryptHash(payload.Password)
	if err != nil {
		return util.Response{
			Ticket:     payload.ID,
			StatusCode: fiber.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = uc.repo.Create(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			return response
		}

		return util.Response{
			Ticket:     payload.ID,
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		}
	}

	// return util.MakeResponse(payload.ID, 200, "Create user successfuly", nil, nil)
	return util.Response{
		Ticket:     payload.ID,
		StatusCode: fiber.StatusOK,
		Message:    "Create user successfuly",
	}

}
