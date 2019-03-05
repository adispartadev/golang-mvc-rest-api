package router

import (
	"golang-mvc-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func SetOwnerRoutes(e *echo.Echo) {
	e.GET("/owners", controller.GetOwners)
}
