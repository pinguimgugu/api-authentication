package resource

import (
	"app/domain"
	"net/http"

	"github.com/labstack/echo"
)

// AuthenticatePostUser struct
type AuthenticatePostUser struct {
	UserRepository domain.UserRepository
}

// Handle method
func (a *AuthenticatePostUser) Handler(c echo.Context) error {
	user := new(domain.RequestAuthenticateUser)

	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	userMap := a.UserRepository.GetByUserAndPassword(user.User, user.Password)

	if _, ok := userMap["username"]; !ok {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, userMap)
}
