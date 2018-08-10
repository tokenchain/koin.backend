package auth

import (
	"github.com/koin-bet/koin.backend/pkg/db"
	"fmt"
)

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
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	return len(exist) > 0 && err == nil
}
