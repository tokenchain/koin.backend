package auth

import (
	"github.com/koin-bet/koin.backend/pkg/db"
	"log"
	"os"
)

var l = log.New(os.Stdout, "[AUTH] ", 0)

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
	u := &struct {
		Hash string `json:"hash"`
	}{}
	_, err := db.GetUser(hash, u)
	if err != nil {
		l.Printf("Error on login: %s\n", err.Error())
		return false
	}
	l.Printf("Structure |%#v|\n", u)
	return u.Hash == hash
}
