package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteTypeHandler struct {
	uc domain.WasteTypeUsecase
	v  *validator.Validate
}

func NewWasteTypeHandler(uc domain.WasteTypeUsecase, v *validator.Validate) domain.WasteTypeHandler {
	return &wasteTypeHandler{
		uc: uc,
		v:  v,
	}
}

func (w *wasteTypeHandler) Create(c fiber.Ctx) error {
	var payload model.WasteTypeCreateOrUpdateRequest

	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	resp := w.uc.Create(c.Context(), payload)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (w *wasteTypeHandler) UpdateByID(c fiber.Ctx) error {
	id := c.Params("id")
	var payload model.WasteTypeCreateOrUpdateRequest

	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	payload.ID = id
	resp := w.uc.UpdateByID(c.Context(), payload)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (w *wasteTypeHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")

	resp := w.uc.DeleteByID(c.Context(), id)
	return c.Status(resp.StatusCode).JSON(resp)
}

func (w *wasteTypeHandler) Find(c fiber.Ctx) error {
	var page int = util.StrToInt(c.Query("page"), 1)
	var limit int = util.StrToInt(c.Query("limit"), 10)

	queryParams := util.NewQueryParams(
		page,
		limit,
		[]types.Filter{},
		types.Order{
			Field: "created_at",
			Dir:   "desc",
		},
	)

	resp := w.uc.Find(c.Context(), queryParams)

	return c.Status(resp.StatusCode).JSON(resp)
}
