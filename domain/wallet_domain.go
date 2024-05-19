package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type WalletHandler interface {
	Create(c fiber.Ctx) error
}

type WalletUsecase interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response
}

type WalletRepository interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) error
}
