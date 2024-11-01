package main

import (
	"smol/controllers"
	"smol/core"
	"smol/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config := core.LoadConfig()
	db := db.ConnectDB(&config)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	app := core.App{
		DB:     db,
		Config: config,
	}

	controllers.InitControllers(e, &app)

	e.Logger.Fatal(e.Start(":8080"))
}
