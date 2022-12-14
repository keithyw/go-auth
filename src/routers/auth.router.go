package routers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/keithyw/auth/controllers"
)

type AuthRouter struct {
	router *mux.Router
}

func NewAuthRouter(controller controllers.AuthController) AuthRouter {
	router := mux.NewRouter().StrictSlash(true)
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/authenticate", controller.AuthenticatePostHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/validate", controller.Validate).Methods(http.MethodPost)
	return AuthRouter{router}
}

func (r AuthRouter) GetRouter() *mux.Router {
	return r.router
}