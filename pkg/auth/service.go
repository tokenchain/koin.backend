package auth

import "../db"

// AuthService check if an user is recognized
type AuthService interface {
	Auth(hash string) bool
}

// Service is a stateless structure that implement AuthService
type Service struct{}

// Auth check if in the database user with this hash exist
func (Service) Auth(hash string) bool {
	exist, err := db.GetDb().HKeys("user." + hash)
	return len(exist) > 0 && err == nil
}
