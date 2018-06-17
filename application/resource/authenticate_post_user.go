package resource

import (
	"app/domain"
	"net/http"

	"github.com/labstack/echo"
)

// AuthenticatePostUser struct
type AuthenticatePostUser struct {
	AuthenticationService domain.AuthenticationService
}

// Handle method
func (a *AuthenticatePostUser) Handler(c echo.Context) error {
	user := new(domain.RequestAuthenticateUser)

	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userMap, err := a.AuthenticationService.Do(user)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"description": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userMap)
}
