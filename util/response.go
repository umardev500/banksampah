package util

import (
	"github.com/google/uuid"
	"github.com/umardev500/banksampah/constant"
	"github.com/umardev500/banksampah/types"
)

type ResponseError struct {
	Code    constant.ErrCodeName `json:"code,omitempty"`
	Details interface{}          `json:"details,omitempty"`
}

type Response struct {
	Ticket     uuid.UUID         `json:"ticket,omitempty"`
	StatusCode int               `json:"-"`
	Message    string            `json:"message,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
	Error      *ResponseError    `json:"error,omitempty"`
	Pagination *types.Pagination `json:"pagination,omitempty"`
}

func MakeResponse(ticket uuid.UUID, code int, message string, data interface{}, err *ResponseError) Response {
	return Response{
		Ticket:     ticket,
		StatusCode: code,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}
