package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/inject"
)

type Api struct {
	app       *fiber.App
	mongoConn *config.MongoConfig
	v         *validator.Validate
}

func New(app *fiber.App, mongoConn *config.MongoConfig, v *validator.Validate) *Api {
	return &Api{
		app:       app,
		mongoConn: mongoConn,
		v:         v,
	}
}

func (api *Api) Register() {
	apiRoute := api.app.Group("/api")
	inject.UserInject(apiRoute, api.mongoConn, api.v)
}
