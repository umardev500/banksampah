package inject

import (
	"github.com/umardev500/banksampah/app/handler"
	"github.com/umardev500/banksampah/app/repository"
	"github.com/umardev500/banksampah/app/usecase"
)

func WasteDepoInject(inject Inject) {
	repo := repository.NewWasteDepoRepository(inject.PgxConfig)
	uc := usecase.NewWasteDepoUsecase(repo)
	handler := handler.NewWasteDepoHandler(uc, inject.V)

	router := inject.Router.Group("/deposits")
	router.Post("/", handler.Deposit)
}
