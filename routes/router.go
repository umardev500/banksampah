package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/routes/api"
)

type Router struct {
	app       *fiber.App
	v         *validator.Validate
	pgxConfig *config.PgxConfig
}

func NewRouter(app *fiber.App, v *validator.Validate, pgxConfig *config.PgxConfig) *Router {
	return &Router{
		app:       app,
		v:         v,
		pgxConfig: pgxConfig,
	}
}

func (r *Router) Register() {
	api.New(r.app, r.v, r.pgxConfig).Register()
}
