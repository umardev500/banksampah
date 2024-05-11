package util

import "github.com/google/uuid"

type Response struct {
	Ticket  uuid.UUID   `json:"ticket,omitempty"`
	Code    int         `json:"-"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func MakeResponse(ticket uuid.UUID, code int, message string, data interface{}) Response {
	return Response{
		Ticket:  ticket,
		Code:    code,
		Message: message,
		Data:    data,
	}
}
