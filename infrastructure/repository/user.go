package repository

import (
	"app/infrastructure/mongodb"

	"gopkg.in/mgo.v2/bson"
)

// User repository using mongodb
type User struct {
	Connection *mongodb.Connection
}

// GetByUserAndPassword method
func (u *User) GetByUserAndPassword(user, password string) map[string]string {
	userMap := make(map[string]string)

	u.Connection.Db.C("users").
		Find(
			bson.M{
				"username": user,
				"password": password,
			}).
		One(&userMap)

	return userMap
}
