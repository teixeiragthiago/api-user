package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/teixeiragthiago/api-user/internal/controller/user"
	"github.com/teixeiragthiago/api-user/internal/middleware"
)

func RegisterUserRoutes(r *gin.Engine, userController *usercontroller.UserController, authMiddleware middleware.AuthenticationMiddleware) {

	users := r.Group("/user")
	{
		users.POST("/", userController.RegisterUser)
		users.POST("/login", userController.Login)
		users.GET("/:id", authMiddleware.ValidateJWT(), userController.GetById)
		users.GET("/", authMiddleware.ValidateJWT(), userController.Get)
		users.PUT("/", authMiddleware.ValidateJWT(), userController.Update)
		users.DELETE("/:id", authMiddleware.ValidateJWT(), userController.Delete)
	}
}
