package routes

import (
	"app/application/resource"
	"app/infrastructure/factory"

	"github.com/labstack/echo"
)

type Authentication struct {
	echo *echo.Echo
}

func NewAuthentication(e *echo.Echo) *Authentication {
	return &Authentication{e}
}

func (a *Authentication) Handler() {

	a.echo.POST("/authenticate/user/", func(c echo.Context) error {
		action := resource.AuthenticatePostUser{factory.GetUserRepository()}

		return action.Handler(c)
	})
}
