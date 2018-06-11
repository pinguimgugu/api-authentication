package domain

// UserRepository interface
type UserRepository interface {
	GetByUserAndPassword(string, string) map[string]string
}
