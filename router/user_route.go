package router

import (
	"golang-mvc-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func SetUserRoutes(e *echo.Echo) {
	e.POST("/users/signup", controller.SignUp)
	e.POST("/users/login", controller.Login)
	e.POST("/users/refresh", controller.RefreshToken)
}
