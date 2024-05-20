package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type wasteDepoHandler struct {
	uc domain.WasteDepoUsecase
	v  *validator.Validate
}

func NewWasteDepoHandler(uc domain.WasteDepoUsecase, v *validator.Validate) domain.WasteDepoHandler {
	return &wasteDepoHandler{
		uc: uc,
		v:  v,
	}
}

func (handler *wasteDepoHandler) ConfirmDeposit(c fiber.Ctx) error {
	id := c.Params("id")

	var payload model.WasteDepoConfirmRequest
	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	// Handle validation
	if hndl, err := util.ValidateJson(c, handler.v, payload); err != nil {
		return hndl
	}

	payload.ID = id
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := handler.uc.ConfirmDeposit(ctx, payload)

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *wasteDepoHandler) Deposit(c fiber.Ctx) (err error) {
	var payload model.WasteDepoCreateRequest
	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	// Handle validation
	if hndl, err := util.ValidateJson(c, handler.v, payload); err != nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := handler.uc.Deposit(ctx, payload)

	return c.Status(resp.StatusCode).JSON(resp)
}
