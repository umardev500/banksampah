package inject

import (
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
)

func WalletInject(inject Inject) {
	repo := repository.NewWalletRepository(inject.PgxConfig)
	uc := usecase.NewWalletUsecase(repo)
	handler := handler.NewWalletHandler(uc, inject.V)

	router := inject.Router.Group("/wallets")
	router.Post("/", handler.Create)
}
