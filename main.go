package main

import (
	"app/application/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	routes.NewAuthentication(e).Handler()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
