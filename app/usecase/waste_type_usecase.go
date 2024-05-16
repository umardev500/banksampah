package usecase

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteTypeUc struct {
	repo domain.WasteTypeRepository
}

func NewWasteTypeUsecase(repo domain.WasteTypeRepository) domain.WasteTypeUsecase {
	return &wasteTypeUc{
		repo: repo,
	}
}

func (uc *wasteTypeUc) Find(ctx context.Context, params *types.QueryParam) util.Response {
	ticket := uuid.New()

	wasteTypes, err := uc.repo.Find(ctx, params)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FailedGetAll))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FailedGetAll))

		return util.Response{
			Ticket:     ticket,
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		}
	}

	var pagination *types.Pagination
	total := params.Pagination.Total

	if total > 0 {
		pagination = &params.Pagination
		pagination.Total = len(wasteTypes)
		rowTotal := 8
		pageTotal := math.Ceil(float64(rowTotal) / float64(pagination.Limit))
		pagination.PageTotal = int(pageTotal)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.WasteType.SuccessGetAll,
		Data:       wasteTypes,
		Pagination: pagination,
	}
}
