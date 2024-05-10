package routes

import (
	"github.com/gorilla/mux"
	"github.com/teixeiragthiago/api-user/internal/controller"
)

func SetupRoutes(router *mux.Router, userController *controller.UserController) {
	router.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/user", userController.GetById).Methods("GET")
}
