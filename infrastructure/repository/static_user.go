package repository

import (
	"app/domain"
	"errors"
)

// StaticUser repository using mongodb
type StaticUser struct {
}

// GetStaticUserRepository method
func GetStaticUserRepository() *StaticUser {
	return &StaticUser{}
}

// GetByUserAndPassword method
func (u *StaticUser) GetByUserAndPassword(user, password string) (*domain.LoggedUser, error) {
	userMap := new(domain.LoggedUser)

	static := map[string]string{
		"user":     "static_example",
		"password": "static_example",
	}

	if static["user"] != user {
		return nil, errors.New("fail")
	}

	userMap.ID = 2
	return userMap, nil
}
