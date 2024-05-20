package types

import "errors"

type Errors error

var (
	ErrNotAffected Errors = errors.New("not affected")
)
