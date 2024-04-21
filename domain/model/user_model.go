package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	Username  string     `json:"username" db:"username"`
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type CreateUser struct {
	Email    string `json:"email" validate:"required,email,min=7"`
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
}
