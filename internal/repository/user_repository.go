package repository

import (
	"github.com/teixeiragthiago/api-user/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (*entity.User, error)
	Get(search string) ([]*entity.User, error)
	Save(user *entity.User) error
	Delete(user *entity.User) error
	Update(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Get(search string) ([]*entity.User, error) {

	var users []*entity.User
	err := r.db.Where("name LIKE '%?%' OR nick LIKE '%?%'", search).Order("name asc").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (r *userRepository) GetById(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Delete(user *entity.User) error {
	err := r.db.Delete(&user, user.ID).Error
	if err != nil {
		return err
	}

	return nil
}

// Update user's data in database
func (r *userRepository) Update(user *entity.User) error {

	err := r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Save(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
