package middleware

import (
	"github.com/labstack/echo/v4"
)

func isAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		value := c.Request().Header.Get("Authorization")

	}
}
