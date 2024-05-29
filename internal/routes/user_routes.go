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
