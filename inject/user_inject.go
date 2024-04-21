package inject

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
	"github.com/umardev500/banksampah/config"
)

func UserInject(router fiber.Router, mongoConn *config.MongoConfig, v *validator.Validate) {
	client := mongoConn.Client

	repo := repository.NewUserRepo()
	uc := usecase.NewUserUsecase(repo, client)
	handler := handler.NewUserHandler(uc, v)

	user := router.Group("/user")

	user.Post("/", handler.Create)
}
