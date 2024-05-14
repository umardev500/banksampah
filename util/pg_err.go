package util

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/umardev500/banksampah/constant"
	"github.com/umardev500/banksampah/types"
)

func GetPgError(errs error) (response Response, err error) {
	if pgErr, ok := errs.(*pgconn.PgError); ok {
		errCode := pgErr.Code
		ticket := uuid.New()
		code := fiber.StatusInternalServerError
		msg := pgErr.Detail
		var clientCode string
		var details interface{}

		// Selecting error code
		switch errCode {
		case string(constant.SqlStateDuplicate):
			// case for duplicate
			code = fiber.StatusBadRequest
			msg = "Duplicate entry detected. Please try again."
			detailMsg, matches := RegexDuplicate(pgErr.Detail)
			details = &types.SqlDuplicateDetail{
				Field: matches[1],
				Value: matches[2],
				Error: detailMsg,
			}
			clientCode = string(constant.ErrCodeNameDuplicate)
		}

		return Response{
			StatusCode: code,
			Ticket:     ticket,
			Message:    msg,
			Error: &ResponseError{
				Code:    constant.ErrCodeName(clientCode),
				Details: details,
			},
		}, errs
	}

	return Response{}, nil
}
