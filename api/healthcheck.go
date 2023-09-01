package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scott-dn/go-boilerplate/internal/app"
)

type healthcheck struct {
	isClosing bool
}

func newHealthCheckService() *healthcheck {
	return &healthcheck{}
}

func (h *healthcheck) Close() {
	h.isClosing = true
}

func registerHealthcheck(server *echo.Echo, app *app.App, hc *healthcheck) {
	server.GET("/health", func(c echo.Context) error {
		if hc.isClosing {
			return echo.ErrServiceUnavailable
		}
		if err := app.HealthCheck(); err != nil {
			return echo.ErrServiceUnavailable
		}
		return c.NoContent(http.StatusOK)
	})
}
