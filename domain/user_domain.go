package domain

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type UserHandler interface {
	Create(c fiber.Ctx) error
}

type UserUsecase interface {
	Create(ctx context.Context, payload model.CreateUser) util.Response
}

type UserRepository interface{}
