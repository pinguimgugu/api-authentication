package resource

import (
	"app/domain"
	"net/http"

	"github.com/labstack/echo"
)

// AuthenticatePostUser struct
type AuthenticatePostUser struct {
}

// Handle method
func (a *AuthenticatePostUser) Handler(c echo.Context) error {
	user := new(domain.RequestAuthenticateUser)

	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, user)
}
