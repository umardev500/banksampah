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

func (uc *walletUsecase) DeleteByID(ctx context.Context, id string) util.Response {
	ticket := uuid.New()
	handler, err := util.ParseIDWithResponse(&id)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedDelete))
		handler.Ticket = ticket
		return *handler
	}

	userID := types.DummyUserID // replace with actual user id
	err = uc.repo.DeleteByID(ctx, id, userID)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedDelete))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedDelete))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Wallet.SuccessDelete,
	}
}

func (uc *walletUsecase) Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response {
	ticket := uuid.New()
	payload.ID = uuid.New().String()
	payload.Type = string(types.WalletExtension)

	result, err := uc.repo.Create(ctx, payload)
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
		Data:       result,
	}
}
