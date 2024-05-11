package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/inject"
)

type Api struct {
	app       *fiber.App
	v         *validator.Validate
	pgxConfig *config.PgxConfig
}

func New(app *fiber.App, v *validator.Validate, pgxConfig *config.PgxConfig) *Api {
	return &Api{
		app:       app,
		v:         v,
		pgxConfig: pgxConfig,
	}
}

func (api *Api) Register() {
	apiRoute := api.app.Group("/api")
	inject.UserInject(apiRoute, api.v, api.pgxConfig)
}
