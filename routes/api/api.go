package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/inject"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct {
	app     *fiber.App
	mongoDB *mongo.Database
	v       *validator.Validate
}

func New(app *fiber.App, mongoDB *mongo.Database, v *validator.Validate) *Api {
	return &Api{
		app:     app,
		mongoDB: mongoDB,
	}
}

func (api *Api) Register() {
	apiRoute := api.app.Group("/api")
	inject.UserInject(apiRoute, api.mongoDB, api.v)
}
