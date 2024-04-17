package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/routes/api"
)

type Router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		app: app,
	}
}

func (r *Router) Register() {
	api.New(r.app).Register()
}
