package domain

import (
	"github.com/gofiber/fiber/v3"
)

type WasteTypeHandler interface {
	Find(c fiber.Ctx) error
}
