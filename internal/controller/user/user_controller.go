package usercontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var userDTO dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	token, err := c.userService.RegisterUser(&userDTO)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	c.httpResponse.Success(w, http.StatusCreated, token)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {

	var userDTO dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	success, err := c.userService.Update(&userDTO)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	c.httpResponse.Success(w, http.StatusOK, success)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {

	idParam := mux.Vars(r)["id"]

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, errors.New("invalid user id"))
	}

	success, err := c.userService.Delete(uint(id))
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	c.httpResponse.Success(w, http.StatusNoContent, success)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
	}

	userResponse, err := c.userService.GetById(uint(id))
	if err != nil {
		http.Error(w, "Error trying to get User", http.StatusBadRequest)
		return
	}

	c.httpResponse.Success(w, http.StatusOK, userResponse)
}

func (c *UserController) Get(w http.ResponseWriter, r *http.Request) {

	searchParam := r.URL.Query().Get("search")

	users, err := c.userService.Get(searchParam)
	if err != nil {
		c.httpResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	c.httpResponse.Success(w, http.StatusOK, users)
}
