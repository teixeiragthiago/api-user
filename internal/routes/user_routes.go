package routes

import (
	"github.com/gorilla/mux"
	usercontroller "github.com/teixeiragthiago/api-user/internal/controller/user"
)

func RegisterUserRoutes(router *mux.Router, userController *usercontroller.UserController) {
	router.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	router.HandleFunc("/user/{id}", userController.GetById).Methods("GET")
	router.HandleFunc("/user", userController.Get).Methods("GET")
	router.HandleFunc("/user/{id}", userController.Delete).Methods("DELETE")
	router.HandleFunc("/user", userController.Update).Methods("PUT")
}
