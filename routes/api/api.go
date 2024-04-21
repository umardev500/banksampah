package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/inject"
)

type Api struct {
	app *fiber.App
	v   *validator.Validate
}

func New(app *fiber.App, v *validator.Validate) *Api {
	return &Api{
		app: app,
		v:   v,
	}
}

func (api *Api) Register() {
	apiRoute := api.app.Group("/api")
	inject.UserInject(apiRoute, api.v)
}
