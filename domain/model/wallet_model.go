package model

type WalletCreateOrUpdateRequest struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
