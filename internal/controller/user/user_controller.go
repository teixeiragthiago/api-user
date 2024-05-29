package usercontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/service"
	"github.com/teixeiragthiago/api-user/internal/util"
)

type UserController struct {
	userService  service.UserService
	httpResponse util.HttpResponseErrorHandler
}

func NewUserController(userService service.UserService, httpResponse util.HttpResponseErrorHandler) *UserController {
	return &UserController{userService, httpResponse}
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
			"error": err,
		})
		return
	}

	token, err := uc.userService.RegisterUser(&userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

	//Utilizando mux

	// err := json.NewDecoder(r.Body).Decode(&userDTO)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// token, err := c.userService.RegisterUser(&userDTO)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// c.httpResponse.Success(w, http.StatusCreated, token)
}

func (uc *UserController) Update(c *gin.Context) {

	var userDTO dto.UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	success, err := uc.userService.Update(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": success,
	})

	// err := json.NewDecoder(r.Body).Decode(&userDTO)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// success, err := c.userService.Update(&userDTO)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// c.httpResponse.Success(w, http.StatusOK, success)
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
		c.JSON(http.StatusOK, gin.H{
			"data": success,
		})
	}

	// idParam := mux.Vars(r)["id"]

	// id, err := strconv.ParseUint(idParam, 10, 64)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, errors.New("invalid user id"))
	// }

	// success, err := c.userService.Delete(uint(id))
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// c.httpResponse.Success(w, http.StatusNoContent, success)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		return
	}

	c.JSON(http.StatusOK, userResponseDTO)

	// idParam := mux.Vars(r)["id"]

	// id, err := strconv.ParseUint(idParam, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid User ID", http.StatusBadRequest)
	// }

	// userResponse, err := c.userService.GetById(uint(id))
	// if err != nil {
	// 	http.Error(w, "Error trying to get User", http.StatusBadRequest)
	// 	return
	// }

	// c.httpResponse.Success(w, http.StatusOK, userResponse)
}

func (uc *UserController) Get(c *gin.Context) {

	searchParam := c.Query("search")

	users, err := uc.userService.Get(searchParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)

	// searchParam := r.URL.Query().Get("search")

	// users, err := c.userService.Get(searchParam)
	// if err != nil {
	// 	c.httpResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	// c.httpResponse.Success(w, http.StatusOK, users)
}
