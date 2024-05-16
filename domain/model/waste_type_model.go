package model

import (
	"time"

	"github.com/google/uuid"
)

type WasteType struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Point       float64    `json:"point"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
