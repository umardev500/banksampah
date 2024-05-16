package inject

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
	"github.com/umardev500/banksampah/config"
)

func WasteTypeInject(router fiber.Router, v *validator.Validate, pgxConfig *config.PgxConfig) {
	repo := repository.NewWasteTypeRepo(pgxConfig)
	uc := usecase.NewWasteTypeUsecase(repo)
	handler := handler.NewWasteTypeHandler(uc, v)

	wasteType := router.Group("/waste-types")

	wasteType.Get("/", handler.Find)
}
