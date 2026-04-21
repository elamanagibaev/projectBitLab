package handlers

import (
	"encoding/json"
	"net/http"
	"projectBit/internal/services"
	"projectBit/models"
)

type UserHandler interface {
	AddUserHandler(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (userHandler *userHandler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		return
	}
	userHandler.userService.AddUserService(newUser)
}
