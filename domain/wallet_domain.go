package domain

import (
	"context"

	"github.com/umardev500/banksampah/domain/model"
	"github.com/umardev500/banksampah/util"
)

type WalletUsecase interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) util.Response
}

type WalletRepository interface {
	Create(ctx context.Context, payload model.WalletCreateOrUpdateRequest) error
}
