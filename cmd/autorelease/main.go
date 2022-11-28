package main

import (
	"github.com/SamirMarin/autorelease/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()

	e.GET("/health", handlers.Health)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
