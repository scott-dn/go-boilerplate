package api

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func initMetrics(server *echo.Echo) {
	// TODO adds middleware to gather metrics
	server.Use(echoprometheus.NewMiddleware("metrics"))
	server.GET("/metrics", echoprometheus.NewHandler())
}
