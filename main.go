package main

import (
	"golang-mvc-rest-api/config"
	r "golang-mvc-rest-api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// setup env
	config.EnvSetup()

	e := echo.New()

	r.InitRoutes(e)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))

}
