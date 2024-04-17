package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type userH struct {
	uc domain.UserUsecase
	v  *validator.Validate
}

func NewUserHandler(uc domain.UserUsecase, v *validator.Validate) domain.UserHandler {
	return &userH{
		uc: uc,
		v:  v,
	}
}

func (uh *userH) Create(c fiber.Ctx) error {
	var payload model.CreateUser

	// Bing body
	if err := c.Bind().Body(&payload); err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	// Handle validation
	if hndl, err := util.ValidateJson(c, uh.v, payload); err != nil {
		return hndl
	}

	return c.JSON("success")
}
