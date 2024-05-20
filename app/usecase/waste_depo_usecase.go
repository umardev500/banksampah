package usecase

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteDepoUsecase struct {
	repo domain.WasteDepoRepository
}

func NewWasteDepoUsecase(repo domain.WasteDepoRepository) domain.WasteDepoUsecase {
	return &wasteDepoUsecase{
		repo: repo,
	}
}

func (uc *wasteDepoUsecase) Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) (resp util.Response) {
	ticket := uuid.New()
	payload.ID = uuid.New().String()

	err := uc.repo.Deposit(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedCreate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedCreate))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Wallet.SuccessCreate,
	}
}
