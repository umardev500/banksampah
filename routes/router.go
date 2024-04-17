package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/routes/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	app     *fiber.App
	mongoDB *mongo.Database
	v       *validator.Validate
}

func NewRouter(app *fiber.App, mongoDB *mongo.Database, v *validator.Validate) *Router {
	return &Router{
		app:     app,
		mongoDB: mongoDB,
		v:       v,
	}
}

func (r *Router) Register() {
	api.New(r.app, r.mongoDB, r.v).Register()
}
