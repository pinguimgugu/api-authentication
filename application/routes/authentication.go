package routes

import (
	"app/application/resource"
	"app/domain"

	"github.com/labstack/echo"
)

// Authentication struct
type Authentication struct {
	echo                  *echo.Echo
	authenticationService domain.AuthenticationService
}

//NewAuthentication func
func NewAuthentication(e *echo.Echo, authService domain.AuthenticationService) *Authentication {
	return &Authentication{e, authService}
}

// Handler
func (a *Authentication) Handler() {

	a.echo.POST("/authenticate/user/", func(c echo.Context) error {
		action := resource.AuthenticatePostUser{a.authenticationService}

		return action.Handler(c)
	})
}
