package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/teixeiragthiago/api-user/internal/dto"
	"github.com/teixeiragthiago/api-user/internal/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.userService.RegisterUser(&userDTO)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created with success!"))
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
	}

	userResponse, err := c.userService.GetById(uint(id))
	if err != nil {
		http.Error(w, "Error trying to get User", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(userResponse)
	if err != nil {
		http.Error(w, "Error encoding/marshing response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
