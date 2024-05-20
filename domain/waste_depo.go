package domain

import (
	"context"

	"github.com/umardev500/banksampah/domain/model"
)

type WasteDepoRepo interface {
	Deposit(ctx context.Context, payload model.WasteDepoCreateRequest)
}
