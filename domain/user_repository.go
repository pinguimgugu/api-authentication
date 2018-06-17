package domain

// UserRepository interface
type UserRepository interface {
	GetByUserAndPassword(string, string) (*LoggedUser, error)
}
