package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type WasteTypeHandler interface {
	Find(c fiber.Ctx) error
}

type WasteTypeUsecase interface {
	Find(ctx context.Context, params *types.QueryParam) util.Response
}

type WasteTypeRepository interface {
	Find(ctx context.Context, params *types.QueryParam) ([]model.WasteType, error)
}
