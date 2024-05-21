package usecase

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteDepoUsecase struct {
	repo          domain.WasteDepoRepository
	walletRepo    domain.WalletRepository
	wasteTypeRepo domain.WasteTypeRepository
	pgxConfig     *config.PgxConfig
}

func NewWasteDepoUsecase(
	repo domain.WasteDepoRepository,
	walletRepo domain.WalletRepository,
	wasteTypeRepo domain.WasteTypeRepository,
	pgxConfig *config.PgxConfig,
) domain.WasteDepoUsecase {
	return &wasteDepoUsecase{
		repo:          repo,
		walletRepo:    walletRepo,
		wasteTypeRepo: wasteTypeRepo,
		pgxConfig:     pgxConfig,
	}
}

func (uc *wasteDepoUsecase) ConfirmDeposit(ctx context.Context, payload model.WasteDepoConfirmRequest) (resp util.Response) {
	ticket := uuid.New()
	payload.Status = model.WasteDepoStatusConfirmed

	handler, err := util.ChekEntireIDFromStructWithResponse(payload)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.FailedParseIDMessage))
		handler.Ticket = ticket
		return *handler
	}

	// wasteTypeID := payload.WasteTypeID
	qty := payload.Quantity
	wtID := payload.WasteTypeID

	var balance *float64
	var wd *model.WasteDepo
	err = uc.pgxConfig.WithTransaction(ctx, func(ctx context.Context) error {
		wd, err = uc.repo.ConfirmDeposit(ctx, payload)
		if err != nil {
			return err
		}
		if qty == 0.0 {
			qty = wd.Quantity
		}
		if wtID == "" {
			wtID = wd.WasteTypeID
		}

		// Find waste category
		wt, err := uc.wasteTypeRepo.FindByID(ctx, wtID)
		if err != nil {
			return err
		}

		// Calculate point
		point := wt.Point * qty

		// Set wallet balance increasing
		var walletBalancePayload model.WalletSetBalanceRequest = model.WalletSetBalanceRequest{
			ID:      wd.WalletID,
			SetType: model.SetIncrease,
			Amount:  point,
		}
		balance, err = uc.walletRepo.SetBalance(ctx, walletBalancePayload)

		return err
	})

	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Deposit.FailedCreate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Deposit.FailedCreate))

		if err == pgx.ErrNoRows {
			return util.NoRowsErrorResponse(ticket)
		}

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Deposit.SuccessCreate,
		Data: map[string]interface{}{
			"wallet_id": wd.WalletID,
			"balance":   balance,
		},
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
		Data: map[string]interface{}{
			"id": payload.ID,
		},
	}
}
