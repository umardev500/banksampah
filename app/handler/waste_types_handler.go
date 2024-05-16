package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/umardev500/banksampah/domain"
	"github.com/umardev500/banksampah/types"
	"github.com/umardev500/banksampah/util"
)

type wasteTypeHandler struct {
	uc domain.WasteTypeUsecase
	v  *validator.Validate
}

func NewWasteTypeHandler(uc domain.WasteTypeUsecase, v *validator.Validate) domain.WasteTypeHandler {
	return &wasteTypeHandler{
		uc: uc,
		v:  v,
	}
}

func (w *wasteTypeHandler) Find(c fiber.Ctx) error {
	var page int = util.StrToInt(c.Query("page"), 10)

	queryParams := util.NewQueryParams(
		page,
		4,
		[]types.Filter{

			{
				Field:    "point",
				Operator: ">=",
				Value:    "800",
			},
		},
		types.Order{
			Field: "created_at",
			Dir:   "desc",
		},
	)

	resp := w.uc.Find(c.Context(), queryParams)

	return c.Status(resp.StatusCode).JSON(resp)
}
