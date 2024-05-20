package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type WasteDepoHandler interface {
	Deposit(c fiber.Ctx) error
}

type WasteDepoUsecase interface {
	Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) util.Response
	ConfirmDeposit(ctx context.Context, payload model.WasteConfirmRequest) util.Response
}

type WasteDepoRepository interface {
	Deposit(ctx context.Context, payload model.WasteDepoCreateRequest) error
	ConfirmDeposit(ctx context.Context, payload model.WasteConfirmRequest) (*model.WasteDepo, error)
	FindByID(ctx context.Context, id string) (*model.WasteDepo, error)
}
