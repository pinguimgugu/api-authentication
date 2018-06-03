package domain

// RequestAuthenticateUser struct
type RequestAuthenticateUser struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
