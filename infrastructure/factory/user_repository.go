package factory

import (
	"app/infrastructure/repository"
	"sync"
)

var onceInstance sync.Once
var userRepositoryOnce *repository.User

// GetUserRepository return instance user repository
func GetUserRepository() *repository.User {

	onceInstance.Do(func() {
		if userRepositoryOnce == nil {
			userRepositoryOnce = &repository.User{Connection: getMongoDb()}
		}
	})

	return userRepositoryOnce
}
