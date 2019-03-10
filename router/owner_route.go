package router

import (
	"fmt"
	"golang-mvc-rest-api/controller"

	"github.com/labstack/echo/v4"
)

func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("======context=======")
		fmt.Printf("%+v\n", c.Param("page"))
		fmt.Println("=============")

		data := 1

		if data != 1 {
			return echo.ErrBadRequest
		}

		return next(c)
	}
}

func isManager(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		data := 2

		if data != 1 {
			return echo.ErrForbidden
		}

		return next(c)
	}

}

func SetOwnerRoutes(e *echo.Echo) {
	// e.GET("/owners", controller.GetOwners)
	e.GET("/owners/page/:page/limit/:limit", controller.GetOwnersLimit)
	e.POST("/owners", controller.AddOwner)
}
