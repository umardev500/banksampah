package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/config"
	"github.com/umardev500/banksampah/routes/api"
)

type Router struct {
	app       *fiber.App
	mongoConn *config.MongoConfig
	v         *validator.Validate
}

func NewRouter(app *fiber.App, mongoConn *config.MongoConfig, v *validator.Validate) *Router {
	return &Router{
		app:       app,
		mongoConn: mongoConn,
		v:         v,
	}
}

func (r *Router) Register() {
	api.New(r.app, r.mongoConn, r.v).Register()
}
