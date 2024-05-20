package types

import "errors"

type Errors error

var (
	ErrNotAffected  Errors = errors.New("not affected")
	ErrIDValidation Errors = errors.New("id validation error")
)

// Struct for key value error detail
type ErrorDetail struct {
	Code  string      `json:"code"`
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Error string      `json:"error"`
}
