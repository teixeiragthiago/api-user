package repository_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserRepository_GetById_Success(t *testing.T) {

	//arrange

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT VERSION()")).WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("1"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := repository.NewUserRepository(gormDB)
	mockUser := &entity.User{
		ID:        1,
		Name:      "Thiago",
		Email:     "thiago@teste.com",
		Nickname:  "thiago_teste",
		Password:  []byte("hashedpassword"),
		CreatedAt: time.Now().Truncate(time.Millisecond),
		Active:    true,
	}
	id := uint(1)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "nickname", "password", "created_at", "active"}).
		AddRow(mockUser.ID, mockUser.Name, mockUser.Email, mockUser.Nickname, mockUser.Password, mockUser.CreatedAt, mockUser.Active)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(id, 1).
		WillReturnRows(rows)

	user, err := userRepo.GetById(id)

	assert.NoError(t, err)

	assert.Equal(t, mockUser, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
