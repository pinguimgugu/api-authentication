package domain

// AuthenticationService interface
type AuthenticationService interface {
	Do(*RequestAuthenticateUser) (*LoggedUser, error)
	AttachRepository(UserRepository)
}
