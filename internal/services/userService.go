package services

import (
	"projectBit/internal/repositories"
	"projectBit/models"
)

type UserService interface {
	AddUserService(user models.User)
}

type userService struct {
	userRepositories repositories.UserRepositories
}

func NewUserService(userRepositories repositories.UserRepositories) UserService {
	return &userService{userRepositories: userRepositories}
}

func (userService *userService) AddUserService(user models.User) {
	userService.userRepositories.AddUser(user)
}
