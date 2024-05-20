package model

import "time"

type WasteDepo struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	WasteTypeID string     `json:"waste_type_id"`
	Quantity    float64    `json:"quantity"`
	Description float64    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type WasteDepoCreateRequest struct {
	ID          string  `json:"-" checkid:"id"`
	UserID      string  `json:"-" checkid:"user_id"`
	WasteTypeID string  `json:"waste_type_id" validate:"required" checkid:"waste_type_id"`
	Quantity    float64 `json:"quantity" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CreatedBy   string  `json:"-" checkid:"created_by"`
}
