package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
