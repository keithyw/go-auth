package repositories

import (
	"github.com/keithyw/auth/models"
)

type AuthRepository interface {
	FindByUsername(username string) (*models.User, error)
}