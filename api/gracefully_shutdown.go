package api

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/scott-dn/go-boilerplate/internal/app"
)

func gracefullyShutdown(app *app.App, server *echo.Echo, hc *healthcheck) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Info().Str("signal", (<-quit).String()).Msg("received signal")

	hc.close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		app.Close()
		cancel()
	}()

	log.Info().Msg("shutting down http server")
	if err := server.Shutdown(ctx); err != nil {
		log.Panic().Err(err).Msg("failed to shutdown http server")
	}
}
