package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with the given details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.UserDTO true "User data"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Error registering user"
// @Router /register [post]
func (uc *UserController) RegisterUser(c *gin.Context) {

	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := uc.userService.RegisterUser(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func (uc *UserController) Update(c *gin.Context) {

	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	success, err := uc.userService.Update(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": success,
	})
}

func (uc *UserController) Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	success, err := uc.userService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"data": success,
		})
	}
}

func (uc *UserController) GetById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	userResponseDTO, err := uc.userService.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error fetching user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponseDTO,
	})
}

func (uc *UserController) Get(c *gin.Context) {

	searchParam := c.Query("search")

	users, err := uc.userService.Get(searchParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
