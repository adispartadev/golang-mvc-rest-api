package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type owner struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Name       string `json:"name"`
}

func GetOwners(c echo.Context) error {
	o := &owner{
		StatusCode: http.StatusBadRequest,
		Message:    "please check your payload",
		Name:       "test",
	}
	return c.JSON(http.StatusBadRequest, o)

}
