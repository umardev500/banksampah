package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type ValidationDetail struct {
	Tag    string `json:"tag"`
	Detail string `json:"detail"`
}

type ValidationResponse struct {
	Message string                      `json:"message"`
	Errors  map[string]ValidationDetail `json:"errors"`
}

func ValidateJson(c fiber.Ctx, v *validator.Validate, payload interface{}) (error, error) {
	errs := make(map[string]ValidationDetail)

	if err := v.Struct(payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				field := strings.ToLower(validationErr.Field())
				tag := validationErr.Tag()

				var detail string = validationMapping(validationErr)
				errs[field] = ValidationDetail{
					Tag:    tag,
					Detail: detail,
				}
			}
		}

		resp := ValidationResponse{
			Message: "Validation error",
			Errors:  errs,
		}

		return c.JSON(resp), errors.New("hi")
	}

	return nil, nil

}

func validationMapping(validationErr validator.FieldError) string {
	var detail string
	field := strings.ToLower(validationErr.Field())
	tag := validationErr.Tag()

	switch tag {
	case "required":
		detail = fmt.Sprintf("Field '%s' is required", field)
	case "min":
		detail = fmt.Sprintf("Field '%s' must be at least %s characters long", field, validationErr.Param())
	case "max":
		detail = fmt.Sprintf("Field '%s' must be at most %s characters long", field, validationErr.Param())
	case "email":
		detail = fmt.Sprintf("Field '%s' must be a valid email address", field)
		// Add more cases for other validation tags as needed
	default:
		detail = fmt.Sprintf("Field %s is %s", field, validationErr.Error())
	}

	return detail
}
