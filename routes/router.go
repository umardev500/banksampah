package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/routes/api"
)

type Router struct {
	app *fiber.App
	v   *validator.Validate
}

func NewRouter(app *fiber.App, v *validator.Validate) *Router {
	return &Router{
		app: app,
		v:   v,
	}
}

func (r *Router) Register() {
	api.New(r.app, r.v).Register()
}
