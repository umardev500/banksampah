package inject

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserInject(router fiber.Router, mongoDB *mongo.Database) {
	repo := repository.NewUserRepo()
	uc := usecase.NewUserUsecase(repo)
	handler := handler.NewUserHandler(uc)

	user := router.Group("/user")

	user.Post("/", handler.Create)
}
