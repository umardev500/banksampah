package util

import (
	"reflect"

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

func CheckIDWithResponse(id string) (resp *Response, err error) {
	_, err = uuid.Parse(id)
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

	return nil, nil
}

func ChekEntireIDFromStruct(src interface{}) []types.ErrorDetail {
	var details []types.ErrorDetail

	v := reflect.ValueOf(src)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		vType := v.Type().Field(i)
		tag, ok := vType.Tag.Lookup("checkid")
		if ok && tag != "" {
			if field.IsZero() {
				details = append(details, types.ErrorDetail{
					Code:  string(constant.ErrCodeNameEmpy),
					Field: tag,
					Error: "The id must not be empty",
				})
				continue
			}

			_, err := uuid.Parse(field.String())
			if err != nil {
				details = append(details, types.ErrorDetail{
					Code:  string(constant.ErrCodeNameInvalidID),
					Field: tag,
					Value: field.String(),
					Error: "The id must be valid UUID.",
				})
				continue
			}
		}
	}

	return details
}

func ChekEntireIDFromStructWithResponse(src interface{}) (resp *Response, err error) {
	details := ChekEntireIDFromStruct(src)
	if len(details) > 0 {
		return &Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    types.InvalidIDMessage,
			Error: &ResponseError{
				Code:    constant.ErrCodeNameInvalidIDs,
				Details: details,
			},
		}, types.ErrIDValidation
	}

	return
}
