package di

import (
	"log"

	"github.com/gorilla/mux"
	pingcontroller "github.com/teixeiragthiago/api-user/internal/controller/ping"
	usercontroller "github.com/teixeiragthiago/api-user/internal/controller/user"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/repository"
	"github.com/teixeiragthiago/api-user/internal/routes"
	"github.com/teixeiragthiago/api-user/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setup dependencies
func setupUserControllerDependencies() (*usercontroller.UserController, error) {
	// Setup database connection, repositories, and services
	db, err := gorm.Open(mysql.Open("connString"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//Migrate the schema
	db.AutoMigrate(&entity.User{})

	//Setup repository
	userRepo := repository.NewUserRepository(db)

	//Setup service
	userService := service.NewUserService(userRepo)

	//Setup controllers
	userController := usercontroller.NewUserController(userService)

	return userController, nil
}

func setupPingControllerDependencies() (*pingcontroller.PingController, error) {
	pingController := pingcontroller.NewPingController()

	return pingController, nil
}

func SetupControllers() *mux.Router {
	router := mux.NewRouter()

	pingController, err := setupPingControllerDependencies()
	if err != nil {
		log.Fatal("Error setting up pingController")
	}

	routes.RegisterPing(router, pingController)

	// userController, err := setupUserControllerDependencies()
	// if err != nil {
	// 	log.Fatal("Error setting up userControllers")
	// }

	// routes.RegisterUserRoutes(router, userController)

	return router
}
