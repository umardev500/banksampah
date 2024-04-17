package domain

import "github.com/gofiber/fiber/v3"

type UserHandler interface {
	Create(c fiber.Ctx) error
}

type UserUsecase interface{}

type UserRepository interface{}
