package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/teixeiragthiago/api-user/internal/controller/user"
)

func RegisterUserRoutes(r *gin.Engine, userController *usercontroller.UserController) {

	users := r.Group("/user")
	{
		users.POST("/", userController.RegisterUser)
		users.GET("/:id", userController.GetById)
		users.GET("/", userController.Get)
		users.PUT("/", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}

// func RegisterUserRoutes(router *mux.Router, userController *usercontroller.UserController) {
// 	router.HandleFunc("/register", userController.RegisterUser).Methods("POST")
// 	router.HandleFunc("/user/{id}", userController.GetById).Methods("GET")
// 	router.HandleFunc("/user", userController.Get).Methods("GET")
// 	router.HandleFunc("/user/{id}", userController.Delete).Methods("DELETE")
// 	router.HandleFunc("/user", userController.Update).Methods("PUT")
// }
