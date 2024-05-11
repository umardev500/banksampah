package util

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/umardev500/banksampah/constant"
)

func GetPgError(errs error) (response Response, err error) {
	if pgErr, ok := errs.(*pgconn.PgError); ok {
		errCode := pgErr.Code
		ticket := uuid.New()
		code := fiber.StatusInternalServerError
		msg := pgErr.Detail

		switch errCode {
		case string(constant.SqlStateDuplicate):
			code = fiber.StatusConflict
			msg = "Conflict, data is already exist"
		}

		return Response{
			Code:    code,
			Ticket:  ticket,
			Message: msg,
		}, errs
	}

	return Response{}, nil
}
