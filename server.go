package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xifanyan/voodoo/database"
	"github.com/xifanyan/voodoo/router"
)

func main() {
	database.Init()

	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initialze router
	router.Init(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
