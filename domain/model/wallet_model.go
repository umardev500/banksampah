package model

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID          uuid.UUID  `json:"id"`
	UserD       uuid.UUID  `json:"user_id,omitempty"`
	Name        string     `json:"name"`
	Balance     float64    `json:"balance"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type WalletCreateOrUpdateRequest struct {
	ID          string `json:"-"`
	UserID      string `json:"-"`
	Name        string `json:"name" validate:"required,min=6"`
	Description string `json:"description"`
	Type        string `json:"-"`
}

type WalletMoveBalanceRequest struct {
	FromWalletID string  `json:"from_wallet_id" validate:"required"`
	ToWalletID   string  `json:"to_wallet_id" validate:"required"`
	Amount       float64 `json:"amount" validate:"required,gt=0"`
}
