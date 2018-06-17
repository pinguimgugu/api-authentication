package service

import (
	"app/domain"
	"errors"
	"sync"
)

//Authentication struct
type Authentication struct {
	userRepositoryList []domain.UserRepository
}

//GetAuthenticationService func
func GetAuthenticationService() *Authentication {
	a := new(Authentication)
	return a
}

// Do method
func (a *Authentication) Do(re *domain.RequestAuthenticateUser) (*domain.LoggedUser, error) {

	wg := sync.WaitGroup{}
	loggedUser := make(chan *domain.LoggedUser)
	endSearch := make(chan bool)

	for key := range a.userRepositoryList {
		wg.Add(1)

		go func(userRepository domain.UserRepository) {
			defer wg.Done()

			user, err := userRepository.GetByUserAndPassword(re.User, re.Password)

			if err == nil {
				loggedUser <- user
			}

		}(a.userRepositoryList[key])
	}

	go func() {
		wg.Wait()
		endSearch <- true
	}()

	select {
	case u := <-loggedUser:
		return u, nil
	case <-endSearch:
		return nil, errors.New("The given username or password not macth an user")
	}
}

// AttachRepository method
func (a *Authentication) AttachRepository(ur domain.UserRepository) {
	a.userRepositoryList = append(a.userRepositoryList, ur)
}
