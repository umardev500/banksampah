package inject

import (
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
)

func WasteDepoInject(inject Inject) {
	repo := repository.NewWasteDepoRepository(inject.PgxConfig)
	walletRepo := repository.NewWalletRepository(inject.PgxConfig)
	wasteTypeRepo := repository.NewWasteTypeRepo(inject.PgxConfig)
	uc := usecase.NewWasteDepoUsecase(repo, walletRepo, wasteTypeRepo, inject.PgxConfig)
	handler := handler.NewWasteDepoHandler(uc, inject.V)

	router := inject.Router.Group("/deposits")
	router.Post("/", handler.Deposit)
	router.Put("/:id/confirm", handler.ConfirmDeposit)
}
