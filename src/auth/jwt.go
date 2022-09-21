package auth

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/keithyw/auth/conf"
)

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateJWT(username string, config conf.Config) (string, error) {
	t := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString([]byte(config.JWTSecretKey))
	return str, err
}

func ValidateJWT(token string, config conf.Config) error {
	t, err := jwt.ParseWithClaims(
		token,
		&JWTClaim{},
		func(tok *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := t.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("could not parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}