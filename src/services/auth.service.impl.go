package services

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/keithyw/auth/models"
	"github.com/keithyw/auth/repositories"
)

type AuthServiceImpl struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &AuthServiceImpl{repo}
}

func (a *AuthServiceImpl) Authenticate(username string, passwd string) (*models.User, error) {
	user, err := a.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
		return nil, err
	}
	return user, nil
}