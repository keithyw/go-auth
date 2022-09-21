package services

import (
	"github.com/keithyw/auth/models"
)

type AuthService interface {
	Authenticate(username string, passwd string) (*models.User, error)
}