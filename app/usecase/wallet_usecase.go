package usecase

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/constant"
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

func (uc *walletUsecase) UpdateByID(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (resp util.Response) {
	ticket := uuid.New()
	handler, err := util.CheckIDWithResponse(payload.ID)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedUpdate))
		handler.Ticket = ticket
		return *handler
	}

	wallet, err := uc.repo.UpdateByID(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedUpdate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FailedUpdate))

		if err == pgx.ErrNoRows {
			return util.Response{
				Ticket:     ticket,
				StatusCode: fiber.StatusBadRequest,
				Message:    types.Wallet.OutOfBalance,
				Error: &util.ResponseError{
					Code:    constant.ErrCodeNameOutOfBalance,
					Details: types.Wallet.OutOfBalanceDetails,
				},
			}
		}

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Wallet.SuccessUpdate,
		Data:       wallet,
	}
}

func (uc *walletUsecase) MoveBalanceToWallet(ctx context.Context, payload model.WalletMoveBalanceRequest) util.Response {
	ticket := uuid.New()
	handler, err := util.CheckIDWithResponse(payload.FromWalletID)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FaildMoveBalance))
		handler.Ticket = ticket
		return *handler
	}

	fromAndToWallet, err := uc.repo.MoveBalanceToWallet(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FaildMoveBalance))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Wallet.FaildMoveBalance))

		if err == pgx.ErrNoRows {
			return util.Response{
				Ticket:     ticket,
				StatusCode: fiber.StatusBadRequest,
				Message:    types.Wallet.OutOfBalance,
				Error: &util.ResponseError{
					Code:    constant.ErrCodeNameOutOfBalance,
					Details: types.Wallet.OutOfBalanceDetails,
				},
			}
		}

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Wallet.SuccessMoveBalance,
		Data:       map[string]interface{}{"from_wallet": fromAndToWallet[0], "to_wallet": fromAndToWallet[1]},
	}
}

func (uc *walletUsecase) DeleteByID(ctx context.Context, id string) util.Response {
	ticket := uuid.New()
	handler, err := util.CheckIDWithResponse(id)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedDelete))
		handler.Ticket = ticket
		return *handler
	}

	err = uc.repo.DeleteByID(ctx, id)
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
