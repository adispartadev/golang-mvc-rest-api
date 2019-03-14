package router

import (
	"golang-mvc-rest-api/controller"
	m "golang-mvc-rest-api/middleware"
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func paramValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[0-9]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(http.StatusBadRequest, "param ("+paramKey[k]+") is not a number")
			}
		}

		return next(c)
	}

}

func SetOwnerRoutes(e *echo.Echo) {
	// e.GET("/owners", controller.GetOwners)
	r := e.Group("/owners", m.IsAuthorized)

	r.GET("/page/:page/limit/:limit", controller.GetOwnersLimit, paramValidation)
	r.POST("/", controller.AddOwner)
	r.DELETE("/:id", controller.RemoveOwner, paramValidation)
	r.PUT("/:id", controller.EditOwner, paramValidation)
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
}
