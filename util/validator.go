package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/umardev500/banksampah/constant"
)

type ValidationDetail struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Error string `json:"error"`
}

func ValidateJson(c fiber.Ctx, v *validator.Validate, payload interface{}) (error, error) {
	// errs := make(map[string]ValidationDetail)
	var errs []ValidationDetail

	if err := v.Struct(payload); err != nil {
		ticket := uuid.New()
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				field := strings.ToLower(validationErr.Field())
				tag := validationErr.Tag()

				switch {
				case tag == "min" || tag == "max":
					tag = fmt.Sprintf("%s=%s", tag, validationErr.Param())
				}

				var errMsg string = validationMapping(validationErr)
				errs = append(errs, ValidationDetail{
					Field: field,
					Tag:   tag,
					Error: errMsg,
				})
			}
		}

		resp := Response{
			Ticket:  ticket,
			Message: "Validation error",
			Error: &ResponseError{
				Code:    constant.ErrCodeNameValidation,
				Details: errs,
			},
		}

		return c.Status(fiber.StatusBadRequest).JSON(resp), err
	}

	return nil, nil

}

func validationMapping(validationErr validator.FieldError) string {
	var detail string
	field := strings.ToLower(validationErr.Field())
	tag := validationErr.Tag()

	switch tag {
	case "required":
		detail = fmt.Sprintf("%s is required.", field)
	case "min":
		detail = fmt.Sprintf("%s must be at least %s characters long.", field, validationErr.Param())
	case "max":
		detail = fmt.Sprintf("%s must be at most %s characters long.", field, validationErr.Param())
	case "email":
		detail = fmt.Sprintf("%s must be a valid email address.", field)
		// Add more cases for other validation tags as needed
	default:
		detail = fmt.Sprintf("Field %s is %s", field, validationErr.Error())
	}

	return detail
}
