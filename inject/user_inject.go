package inject

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
)

func UserInject(router fiber.Router, v *validator.Validate) {

	repo := repository.NewUserRepo()
	uc := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(uc, v)

	user := router.Group("/user")

	user.Post("/", handler.Create)
}
