package main

import (
	"log"
	"net/http"
	"github.com/keithyw/auth/conf"
	"github.com/keithyw/auth/controllers"
	"github.com/keithyw/auth/database"
	"github.com/keithyw/auth/repositories"
	"github.com/keithyw/auth/routers"
	"github.com/keithyw/auth/services"
)

func main() {
	config, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}
	db := database.NewDatabase(config)
	defer db.DB.Close()
	repo := repositories.NewAuthRepository(db)
	service := services.NewAuthService(repo)
	controller := controllers.NewAuthController(service, config)
	r := routers.NewAuthRouter(controller)
	log.Fatal(http.ListenAndServe(config.Port, r.GetRouter()))
}