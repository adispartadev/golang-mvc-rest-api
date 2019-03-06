package main

import (
	"golang-mvc-rest-api/config"
	"golang-mvc-rest-api/db"
	r "golang-mvc-rest-api/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// setup env
	config.EnvSetup()
	db.Connect()

	e := echo.New()

	r.InitRoutes(e)

	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))

}
