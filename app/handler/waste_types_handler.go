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
	var page int = util.StrToInt(c.Query("page"), 1)
	var limit int = util.StrToInt(c.Query("limit"), 10)

	queryParams := util.NewQueryParams(
		page,
		limit,
		[]types.Filter{},
		types.Order{
			Field: "created_at",
			Dir:   "desc",
		},
	)

	resp := w.uc.Find(c.Context(), queryParams)

	return c.Status(resp.StatusCode).JSON(resp)
}
