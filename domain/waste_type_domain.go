package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type WasteTypeHandler interface {
	Create(c fiber.Ctx) error
	Find(c fiber.Ctx) error
	DeleteByID(c fiber.Ctx) error
	UpdateByID(c fiber.Ctx) error
}

type WasteTypeUsecase interface {
	Create(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) util.Response
	Find(ctx context.Context, params *types.QueryParam) util.Response
	DeleteByID(ctx context.Context, id string) util.Response
	UpdateByID(ctx context.Context, pyload model.WasteTypeCreateOrUpdateRequest) util.Response
}

type WasteTypeRepository interface {
	Create(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) (*model.WasteType, error)
	Find(ctx context.Context, params *types.QueryParam) (*model.FindWasteTypeResponse, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) error
}
