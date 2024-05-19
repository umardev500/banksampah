package model

type WalletCreateOrUpdateRequest struct {
	ID          string `json:"-"`
	UserID      string `json:"-"`
	Name        string `json:"name" validate:"required,min=6"`
	Description string `json:"description"`
}
