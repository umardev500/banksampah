package util

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/umardev500/banksampah/constant"
	"github.com/umardev500/banksampah/types"
)

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func NewUUIDPointer() *uuid.UUID {
	uuid := GenerateUUID()
	return &uuid
}

func ParseIDWithHandler(id *string) (resp *Response, err error) {
	idUUID, err := uuid.Parse(*id)
	if err != nil {
		return &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    types.InvalidIDMessage,
			Error: &ResponseError{
				Code:    constant.ErrCodeNameInvalidID,
				Details: types.MustUUIDValidError,
			},
		}, err
	}

	*id = idUUID.String()

	return nil, nil
}
