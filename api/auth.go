package api

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.ErrUnauthorized
			}
			if !strings.HasPrefix(authorization, "Bearer ") {
				return echo.ErrUnauthorized
			}
			authorization = strings.TrimPrefix(authorization, "Bearer ")
			token, err := jwt.ParseWithClaims(authorization, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return app.Config.JwtSecret, nil
			})
			if err != nil {
				return echo.ErrUnauthorized
			}
			claims, ok := token.Claims.(*jwtCustomClaims)
			if !ok || !token.Valid {
				return echo.ErrUnauthorized
			}
			c.Set("email", claims.Email)
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
