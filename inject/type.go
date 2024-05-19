package inject

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/config"
)

type Inject struct {
	Router    fiber.Router
	V         *validator.Validate
	PgxConfig *config.PgxConfig
}
