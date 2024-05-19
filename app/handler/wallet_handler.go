package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type walletHandler struct {
	uc domain.WalletUsecase
	v  *validator.Validate
}

func NewWalletHandler(uc domain.WalletUsecase, v *validator.Validate) domain.WalletHandler {
	return &walletHandler{
		uc: uc,
		v:  v,
	}
}

func (handler *walletHandler) MoveBalanceToWallet(c fiber.Ctx) error {
	var payload model.WalletMoveBalanceRequest
	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	// Handle validation
	if hndl, err := util.ValidateJson(c, handler.v, payload); err != nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := handler.uc.MoveBalanceToWallet(ctx, payload)

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *walletHandler) DeleteByID(c fiber.Ctx) error {
	id := c.Params("id")

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := handler.uc.DeleteByID(ctx, id)

	return c.Status(resp.StatusCode).JSON(resp)
}

func (handler *walletHandler) Create(c fiber.Ctx) error {
	var payload model.WalletCreateOrUpdateRequest
	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	// Handle validation
	if hndl, err := util.ValidateJson(c, handler.v, payload); err != nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	// Append user id
	payload.UserID = types.DummyUserID // replace with actual user id

	resp := handler.uc.Create(ctx, payload)

	return c.Status(resp.StatusCode).JSON(resp)
}
