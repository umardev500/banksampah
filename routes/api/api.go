package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/inject"
)

type Api struct {
	app *fiber.App
}

func New(app *fiber.App) *Api {
	return &Api{
		app,
	}
}

func (a *Api) Register() {
	api := a.app.Group("/api")
	inject.UserInject(api)
}
