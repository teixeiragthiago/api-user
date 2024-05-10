package repository

import (
	"github.com/teixeiragthiago/api-user/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *entity.User) error
	GetById(id uint) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetById(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}
