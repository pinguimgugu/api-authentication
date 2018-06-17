package repository

import (
	"app/domain"
	"app/infrastructure/mongodb"

	"gopkg.in/mgo.v2/bson"
)

// User repository using mongodb
type User struct {
	Connection *mongodb.Connection
}

// GetByUserAndPassword method
func (u *User) GetByUserAndPassword(user, password string) (*domain.LoggedUser, error) {
	userMap := new(domain.LoggedUser)

	err := u.Connection.Db.C("users").
		Find(
			bson.M{
				"username": user,
				"password": password,
			}).
		One(&userMap)

	if err != nil {
		return nil, err
	}

	return userMap, nil
}
