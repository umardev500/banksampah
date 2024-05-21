package model

import "time"

type WasteDepoStatus string

var (
	WasteDepoStatusConfirmed   WasteDepoStatus = "confirmed"
	WasteDepoStatusUnConfirmed WasteDepoStatus = "unconfirmed"
)

type WasteDepo struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	WalletID    string          `json:"wallet_id"`
	WasteTypeID string          `json:"waste_type_id"`
	Quantity    float64         `json:"quantity"`
	Description string          `json:"description"`
	Status      WasteDepoStatus `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	DeletedAt   *time.Time      `json:"deleted_at,omitempty"`
	CreatedBy   *string         `json:"created_by,omitempty"`
	DeletedBy   *string         `json:"deleted_by,omitempty"`
}

type WasteDepoCreateRequest struct {
	ID          string          `json:"-" checkid:"id"`
	UserID      string          `json:"-" checkid:"user_id"`
	WalletID    string          `json:"wallet_id" checkid:"wallet_id"`
	WasteTypeID string          `json:"waste_type_id" validate:"required" checkid:"waste_type_id"`
	Quantity    float64         `json:"quantity" validate:"required"`
	Description string          `json:"description" validate:"required"`
	Status      WasteDepoStatus `json:"-"`
	CreatedBy   string          `json:"-" checkid:"created_by"`
}

type WasteDepoConfirmRequest struct {
	ID          string          `json:"-" checkid:"id"`
	WasteTypeID string          `json:"waste_type_id" db:"waste_type_id"` // optional
	Quantity    float64         `json:"quantity" db:"quantity"`           // optional
	Status      WasteDepoStatus `json:"-" db:"status"`
}

// Delete deposit by id type
type WasteDepoDeleteByIDRequest struct {
	ID        string `json:"-" checkid:"id"`
	DeletedBy string `json:"-" checkid:"deleted_by"`
}
