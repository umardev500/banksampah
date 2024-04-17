package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/routes/api"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	app     *fiber.App
	mongoDB *mongo.Database
}

func NewRouter(app *fiber.App, mongoDB *mongo.Database) *Router {
	return &Router{
		app:     app,
		mongoDB: mongoDB,
	}
}

func (r *Router) Register() {
	api.New(r.app, r.mongoDB).Register()
}
