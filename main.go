package main

import (
	"app/application/routes"
	"app/domain/service"
	"app/infrastructure/factory"
	"app/infrastructure/repository"

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

	authenticationService := service.GetAuthenticationService()
	authenticationService.AttachRepository(factory.GetUserRepository())
	authenticationService.AttachRepository(repository.GetStaticUserRepository())

	routes.NewAuthentication(e, authenticationService).Handler()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
