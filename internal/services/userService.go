package services

import (
	"errors"
	"projectBit/internal/repositories"
	"projectBit/models"
)

type UserService interface {
	AddUser(user models.User) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) AddUser(user models.User) (models.User, error) {
	if user.Name == "" || user.Email == "" {
		return models.User{}, errors.New("name и email обязательны")
	}
	newUser, err := s.userRepository.AddUser(user)

	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}
