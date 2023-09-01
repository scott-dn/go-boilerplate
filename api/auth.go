package api

import (
	"github.com/labstack/echo/v4"
	"github.com/scott-dn/go-boilerplate/internal/app"
)

func auth(app *app.App) echo.MiddlewareFunc {
	// for local development,
	// we send email in 'Authorization' header
	if app.Config.GoENV == "local" {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				authorization := c.Request().Header.Get("Authorization")
				if authorization == "" {
					return echo.ErrUnauthorized
				}
				c.Set("email", authorization)
				return next(c)
			}
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: jwt
			return next(c)
		}
	}
}

func requireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		email, ok := c.Get("email").(string)
		if !ok {
			return echo.ErrForbidden
		}
		if email == "admin@example.com" {
			return next(c)
		}
		return echo.ErrForbidden
	}
}
