package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
)

type wasteTypeHandler struct{}

func NewWasteTypeHandler() domain.WasteTypeHandler {
	return &wasteTypeHandler{}
}

func (w *wasteTypeHandler) Find(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
