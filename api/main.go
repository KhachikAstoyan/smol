package main

import (
	"smol/controllers"
	"smol/core"
	"smol/db"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config := core.LoadConfig()
	db := db.ConnectDB(&config)

	app := core.App{
		DB:     db,
		Config: config,
	}

	controllers.InitControllers(e, &app)

	e.Logger.Fatal(e.Start(":8080"))
}
