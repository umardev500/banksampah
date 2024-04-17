package inject

import (
	"github.com/gofiber/fiber/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserInject(router fiber.Router, mongoDB *mongo.Database) {}
