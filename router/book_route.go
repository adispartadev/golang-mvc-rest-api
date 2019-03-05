package router

import (
	"golang-mvc-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func SetBookRoutes(e *echo.Echo) {
	e.GET("/books", controller.GetBooks)
}
