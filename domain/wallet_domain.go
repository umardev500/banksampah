package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type WalletHandler interface {
	Create(c fiber.Ctx) error
	DeleteByID(c fiber.Ctx) error
	MoveBalanceToWallet(c fiber.Ctx) error
	UpdateByID(c fiber.Ctx) error
}

type WalletUsecase interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response
	DeleteByID(ctx context.Context, id string) util.Response
	MoveBalanceToWallet(ctx context.Context, payload model.WalletMoveBalanceRequest) util.Response
	UpdateByID(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response
}

type WalletRepository interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (*model.Wallet, error)
	DeleteByID(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (model.Wallet, error)
	MoveBalanceToWallet(ctx context.Context, payload model.WalletMoveBalanceRequest) ([]model.Wallet, error)
	UpdateByID(ctx context.Context, payload model.WalletCreateOrUpdateRequest) (*model.Wallet, error)
	SetBalance(ctx context.Context, payload model.WalletSetBalanceRequest) (*float64, error)
}
