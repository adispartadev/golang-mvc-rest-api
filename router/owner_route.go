package router

import (
	"fmt"
	"golang-mvc-rest-api/controller"
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
	e.GET("/owners/page/:page/limit/:limit", controller.GetOwnersLimit, paramValidation)
	e.POST("/owners", controller.AddOwner)
	e.DELETE("/owners/:id", controller.RemoveOwner, paramValidation)
	e.PUT("/owners/:id", controller.EditOwner, paramValidation)
	e.GET("/owners", func(c echo.Context) error {
		req := c.Request()
		fmt.Println("=============")
		fmt.Printf("%+v\n", req.Proto)
		fmt.Printf("%+v\n", req.Host)
		fmt.Printf("%+v\n", req.RemoteAddr)
		fmt.Printf("%+v\n", req.Method)
		fmt.Printf("%+v\n", req.URL.Path)
		fmt.Println("=============")
		return c.JSON(http.StatusOK, "ok")
	})
}
