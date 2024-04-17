package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
)

type userH struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) domain.UserHandler {
	return &userH{
		uc: uc,
	}
}

func (uh *userH) Create(c fiber.Ctx) error {
	var payload model.CreateUser

	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(payload)
}
