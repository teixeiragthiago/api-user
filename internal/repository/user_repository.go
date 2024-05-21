package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/teixeiragthiago/api-user/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id uint) (*entity.User, error)
	Get(search string) ([]*entity.User, error)
	Exists(prop string, input string) (bool, error)
	Save(user *entity.User) error
	Delete(user *entity.User) error
	Update(user *entity.User) error
	InUseByAnotherUser(userID uint, prop string, input string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) InUseByAnotherUser(userID uint, prop string, input string) (bool, error) {

	var user *entity.User

	sqlQuery := fmt.Sprintf("%s = LOWER(?) AND id != ?", prop)
	result := r.db.Where(sqlQuery, input, userID).First(&user)
	return result.RowsAffected > 0, nil
}

func (r *userRepository) Exists(prop string, input string) (bool, error) {

	var user entity.User

	sqlQuery := fmt.Sprintf("%s = LOWER(?)", prop)
	result := r.db.Where(sqlQuery, strings.ToLower(input)).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (r *userRepository) Get(search string) ([]*entity.User, error) {

	var users []*entity.User

	sqlQuery := "(name LIKE LOWER(?) OR nickname LIKE LOWER(?) OR email LIKE LOWER(?)) AND active = true"
	sqlValues := fmt.Sprintf("%%%s%%", search)
	err := r.db.Where(sqlQuery, sqlValues, sqlValues, sqlValues).Order("name ASC").Find(&users).Error
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

func (r *userRepository) Update(user *entity.User) error {

	user.CreatedAt = time.Now()
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
