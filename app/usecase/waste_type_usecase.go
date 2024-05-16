package usecase

import (
	"context"
	"math"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
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

func (uc *wasteTypeUc) UpdateByID(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) util.Response {
	ticket := uuid.New()
	handler, err := util.ParseIDWithResponse(&payload.ID)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FaildUpdate))
		handler.Ticket = ticket
		return *handler
	}

	err = uc.repo.UpdateByID(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FaildUpdate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FaildUpdate))

		return util.Response{
			Ticket:     ticket,
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		}
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.WasteType.SuccessUpdate,
	}
}

func (uc *wasteTypeUc) DeleteByID(ctx context.Context, id string) util.Response {
	ticket := uuid.New()
	handler, err := util.ParseIDWithResponse(&id)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FailedDelete))
		handler.Ticket = ticket
		return *handler
	}

	err = uc.repo.DeleteByID(ctx, id)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FailedDelete))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.WasteType.FailedDelete))

		return util.Response{
			Ticket:     ticket,
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		}
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.WasteType.SuccessDelete,
	}
}

func (uc *wasteTypeUc) Find(ctx context.Context, params *types.QueryParam) util.Response {
	ticket := uuid.New()

	response, err := uc.repo.Find(ctx, params)
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

	if response.Total > 0 {
		pagination = &params.Pagination
		pagination.Total = response.Total
		pageTotal := math.Ceil(float64(response.Total) / float64(pagination.Limit))
		pagination.PageTotal = int(pageTotal)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.WasteType.SuccessGetAll,
		Data:       response.WasteTypes,
		Pagination: pagination,
	}
}
