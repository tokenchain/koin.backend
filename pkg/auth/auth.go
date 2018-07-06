package auth

import "github.com/koinkoin-io/koinkoin.backend/pkg/db"

// AuthService check if an user is recognized
type AuthService interface {
	Auth(hash string) bool
}

// Auth is a stateless structure that implement AuthService
type Auth struct{}

// New return a new service auth
func New() Auth {
	return Auth{}
}

// Auth check if in the database user with this hash exist
func (Auth) Auth(hash string) bool {
	exist, err := db.GetDb().HKeys("user." + hash)
	return len(exist) > 0 && err == nil
}
