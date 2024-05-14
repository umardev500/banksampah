package util

import (
	"github.com/google/uuid"
	"github.com/umardev500/banksampah/constant"
)

type ResponseError struct {
	Code    constant.ErrCodeName `json:"code,omitempty"`
	Details interface{}          `json:"details,omitempty"`
}

type Response struct {
	Ticket  uuid.UUID      `json:"ticket,omitempty"`
	Code    int            `json:"-"`
	Message string         `json:"message,omitempty"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ResponseError `json:"error,omitempty"`
}

func MakeResponse(ticket uuid.UUID, code int, message string, data interface{}) Response {
	return Response{
		Ticket:  ticket,
		Code:    code,
		Message: message,
		Data:    data,
	}
}
