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

type walletUsecase struct {
	repo domain.WalletRepository
}

func NewWalletUsecase(repo domain.WalletRepository) domain.WalletUsecase {
	return &walletUsecase{
		repo: repo,
	}
}

func (uc *walletUsecase) Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response {
	ticket := uuid.New()
	payload.ID = ticket.String()

	err := uc.repo.Create(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FaildUpdate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FaildUpdate))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Wallet.SuccessCreate,
	}
}
