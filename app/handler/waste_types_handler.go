package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
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

func (w *wasteTypeHandler) Find(c fiber.Ctx) error {
	resp := w.uc.Find(c.Context())

	return c.Status(resp.StatusCode).JSON(resp)
}
