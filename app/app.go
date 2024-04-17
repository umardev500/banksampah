package app

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/banksampah/routes"
)

type App struct{}

func New() *App {
	return &App{}
}

func (app *App) Run(ctx context.Context) error {
	fiberApp := fiber.New()

	routes.NewRouter(fiberApp).Register() // register routes

	ch := make(chan error, 1)
	go func() {
		port := os.Getenv("PORT")
		addr := ":" + port

		log.Info().Msgf("Listening on %s 🔥", port)
		ch <- fiberApp.Listen(addr, fiber.ListenConfig{
			DisableStartupMessage: true,
		})
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Println() // add new line
		log.Info().Msgf("Gracefully shutting down... 😪")
		return fiberApp.Shutdown()
	}
}
