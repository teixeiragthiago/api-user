package repository_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/teixeiragthiago/api-user/internal/entity"
	"github.com/teixeiragthiago/api-user/internal/repository"
	testutils "github.com/teixeiragthiago/api-user/internal/test_utils"
)

func TestUserRepository_InUserByAnotherUser_True(t *testing.T) {
	//Arrange
	gormDB, mock := testutils.SetupMockDB(t)
	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	userRepo := repository.NewUserRepository(gormDB)
	userMock := mockValidUser()
	rowsMock := mockRows(userMock)

	userId := uint(1)
	propParameter := "nickname"
	input := "thiago_teste"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE ? = LOWER(?) AND id != ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(propParameter, input, userId, 1).
		WillReturnRows(rowsMock)

	//Act
	alreadyInUse, err := userRepo.InUseByAnotherUser(userId, propParameter, input)

	//Assert
	assert.NoError(t, err)

	assert.Equal(t, true, alreadyInUse)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_GetById_Success(t *testing.T) {

	//Arrange
	gormDB, mock := testutils.SetupMockDB(t)
	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	userRepo := repository.NewUserRepository(gormDB)
	userMock := mockValidUser()
	id := uint(1)

	rows := mockRows(userMock)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(id, 1).
		WillReturnRows(rows)

	//Act
	user, err := userRepo.GetById(id)

	//Assert
	assert.NoError(t, err)

	assert.Equal(t, userMock, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_GetById_Error(t *testing.T) {
	//Arrange
	gormDB, mock := testutils.SetupMockDB(t)
	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	userRepo := repository.NewUserRepository(gormDB)
	id := uint(1)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(id, 1).
		WillReturnError(errors.New("user not found"))

	//Act
	user, err := userRepo.GetById(id)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Nil(t, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_GetByEmail_Success(t *testing.T) {

	//Arrange
	gormDB, mock := testutils.SetupMockDB(t)
	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	userRepo := repository.NewUserRepository(gormDB)
	userMock := mockValidUser()
	email := "thiago@teste.com"

	rows := mockRows(userMock)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(email, 1).
		WillReturnRows(rows)

	//Act
	user, err := userRepo.GetByEmail(email)

	//Assert
	assert.NoError(t, err)

	assert.Equal(t, userMock, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUserRepository_GetByEmail_Error(t *testing.T) {
	//Arrange
	gormDB, mock := testutils.SetupMockDB(t)
	defer func() {
		db, _ := gormDB.DB()
		db.Close()
	}()

	userRepo := repository.NewUserRepository(gormDB)
	email := "thiago@teste.com"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(email, 1).
		WillReturnError(errors.New("user not found"))

	//Act
	user, err := userRepo.GetByEmail(email)

	//Assert
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Nil(t, user)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func mockRows(mockUser *entity.User) *sqlmock.Rows {

	return sqlmock.NewRows([]string{"id", "name", "email", "nickname", "password", "created_at", "active"}).
		AddRow(mockUser.ID, mockUser.Name, mockUser.Email, mockUser.Nickname, mockUser.Password, mockUser.CreatedAt, mockUser.Active)
}

func mockValidUser() *entity.User {
	return &entity.User{
		ID:        1,
		Name:      "Thiago",
		Email:     "thiago@teste.com",
		Nickname:  "thiago_teste",
		Password:  []byte("hashedpassword"),
		CreatedAt: time.Now().Truncate(time.Millisecond),
		Active:    true,
	}
}
