package di

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/teixeiragthiago/api-user/config"
	pingcontroller "github.com/teixeiragthiago/api-user/internal/controller/ping"
	usercontroller "github.com/teixeiragthiago/api-user/internal/controller/user"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/middleware"
	"github.com/teixeiragthiago/api-user/internal/repository"
	"github.com/teixeiragthiago/api-user/internal/routes"
	"github.com/teixeiragthiago/api-user/internal/service"
	"github.com/teixeiragthiago/api-user/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.ConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Setup dependencies
func setupUserControllerDependencies() (*usercontroller.UserController, error) {
	// Setup database connection, repositories, and services
	db, err := setupDatabase()
	if err != nil {
		panic("Erro ao conectar com o banco de dados!")
	}

	//Migrate the schema
	//TODO https://atlasgo.io/guides/orms/gorm
	db.AutoMigrate(&entity.User{}, &entity.Post{})

	//Setup repository
	userRepo := repository.NewUserRepository(db)

	//TokenGenerator
	jwtService := util.NewJWTService(config.JwtKey)

	//Setup service
	userService := service.NewUserService(userRepo, jwtService)

	//Setup controllers
	userController := usercontroller.NewUserController(userService)

	return userController, nil
}

func setupPingControllerDependencies() (*pingcontroller.PingController, error) {
	pingController := pingcontroller.NewPingController()

	return pingController, nil
}

func setupAuthMiddleware() middleware.AuthenticationMiddleware {

	jwtService := util.NewJWTService(config.JwtKey)

	return middleware.NewAuthMiddleware(jwtService)

}

func SetupDependencies() *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	pingController, err := setupPingControllerDependencies()
	if err != nil {
		log.Fatal("Error setting up pingController")
	}

	routes.RegisterPing(router, pingController)

	router.Run(fmt.Sprintf("%d", config.ApiPort))

	userController, err := setupUserControllerDependencies()
	if err != nil {
		log.Fatal("Error setting up userControllers")
	}

	authMiddleware := setupAuthMiddleware()

	routes.RegisterUserRoutes(router, userController, authMiddleware)

	return router
}
