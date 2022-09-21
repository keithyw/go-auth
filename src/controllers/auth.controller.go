package controllers

import (
	"context"
	"fmt"
	"io"
	"strings"
	"encoding/json"
	"net/http"
	"github.com/golang-jwt/jwt"
	"github.com/keithyw/auth/auth"
	"github.com/keithyw/auth/conf"
	"github.com/keithyw/auth/models"
	"github.com/keithyw/auth/services"
)

type AuthController struct {
	Svc services.AuthService
	config *conf.Config
}

type authResponse struct {
	Token string `json:"token,omitempty"`
}

func NewAuthController(svc services.AuthService, config *conf.Config) AuthController {
	return AuthController{svc, config}
}

func (a *AuthController) AuthenticatePostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}
	var u models.User
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	authAuth, err := a.Svc.Authenticate(u.Username, u.Passwd)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	token, err := auth.CreateJWT(authAuth.Username, *a.config)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	jsonString, err := json.Marshal(authResponse{token})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.Write(jsonString)
}

func (a *AuthController) Validate(w http.ResponseWriter, r *http.Request) {
	header := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(header) != 2 {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Malformed token"))
	} else {
		jwtToken := header[1]
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])

			}
			return []byte(a.config.JWTSecretKey), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			_ = context.WithValue(r.Context(), "props", claims)
			w.Header().Set("context-type", "application/json")
			w.Write([]byte("Authorized"))
			// ctx.Value()
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	}
}