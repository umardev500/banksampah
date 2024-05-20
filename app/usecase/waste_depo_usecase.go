package usecase

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteDepoUsecase struct {
	repo       domain.WasteDepoRepository
	walletRepo domain.WalletRepository
	pgxConfig  *config.PgxConfig
}

func NewWasteDepoUsecase(repo domain.WasteDepoRepository, walletRepo domain.WalletRepository, pgxConfig *config.PgxConfig) domain.WasteDepoUsecase {
	return &wasteDepoUsecase{
		repo:       repo,
		walletRepo: walletRepo,
		pgxConfig:  pgxConfig,
	}
}

func (uc *wasteDepoUsecase) Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) (resp util.Response) {
	ticket := uuid.New()
	payload.ID = uuid.New().String()
	payload.UserID = types.DummyUserID     // Replace with actual id
	payload.CreatedBy = types.DummyAdminID // Replace with actual id

	handler, err := util.ChekEntireIDFromStructWithResponse(payload)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.FailedParseIDMessage))
		handler.Ticket = ticket
		return *handler
	}

	err = uc.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		err = uc.repo.Deposit(ctx, payload)
		if err != nil {
			return err
		}

		// Find waste category

		// Set wallet balance increasing
		var walletBalancePayload model.WalletSetBalanceRequest = model.WalletSetBalanceRequest{
			ID:      payload.WasteTypeID,
			SetType: model.SetIncrease,
			Amount:  1000,
		}
		_, err = uc.walletRepo.SetBalance(ctx, walletBalancePayload)

		return err
	})
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Deposit.FailedCreate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Deposit.FailedCreate))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Deposit.SuccessCreate,
	}
}
