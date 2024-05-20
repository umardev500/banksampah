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

func (uc *wasteTypeUc) Create(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) util.Response {
	ticket := uuid.New()
	payload.ID = uuid.New().String()

	result, err := uc.repo.Create(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FaildUpdate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FaildUpdate))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Waste.SuccessCreate,
		Data:       result,
	}
}

func (uc *wasteTypeUc) UpdateByID(ctx context.Context, payload model.WasteTypeCreateOrUpdateRequest) util.Response {
	ticket := uuid.New()
	handler, err := util.CheckIDWithResponse(payload.ID)
	if err != nil {
		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FaildUpdate))
		handler.Ticket = ticket
		return *handler
	}

	err = uc.repo.UpdateByID(ctx, payload)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FaildUpdate))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FaildUpdate))

		return util.InternalErrorResponse(ticket)
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Waste.SuccessUpdate,
	}
}

func (uc *wasteTypeUc) DeleteByID(ctx context.Context, id string) util.Response {
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
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedDelete))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedDelete))

		return util.Response{
			Ticket:     ticket,
			StatusCode: fiber.StatusInternalServerError,
			Message:    fiber.ErrInternalServerError.Message,
		}
	}

	return util.Response{
		Ticket:     ticket,
		StatusCode: fiber.StatusOK,
		Message:    types.Waste.SuccessDelete,
	}
}

func (uc *wasteTypeUc) Find(ctx context.Context, params *types.QueryParam) util.Response {
	ticket := uuid.New()

	response, err := uc.repo.Find(ctx, params)
	if err != nil {
		if response, isPgErr := util.GetPgError(err); isPgErr != nil {
			log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedGetAll))
			return response
		}

		log.Error().Msgf(util.LogParseError(&ticket, err, types.Waste.FailedGetAll))

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
		Message:    types.Waste.SuccessGetAll,
		Data:       response.WasteTypes,
		Pagination: pagination,
	}
}
