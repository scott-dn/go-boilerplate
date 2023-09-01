package api

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func newValidator() *customValidator {
	return &customValidator{validator: validator.New()}
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, validationError := range validationErrors {
				return echo.NewHTTPError(http.StatusBadRequest, validationError)
			}
		}
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}
