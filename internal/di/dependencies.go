package di

import (
	"github.com/teixeiragthiago/api-user/internal/controller"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/repository"
	"github.com/teixeiragthiago/api-user/internal/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Setup dependencies
func SetupDependencies() (*controller.UserController, error) {
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
	userController := controller.NewUserController(userService)

	return userController, nil
}
