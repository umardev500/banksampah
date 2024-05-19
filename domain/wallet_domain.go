package domain

import (
	"context"

	"github.com/umardev500/banksampah/domain/model"
)

type WalletRepository interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) error
}
