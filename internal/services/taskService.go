package services

import (
	"errors"
	"projectBit/internal/repositories"
	"projectBit/models"
)

type TaskService interface {
	AddTask(tasks models.Tasks) (models.Tasks, error)
	ChangeTask(task models.Tasks) error
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) TaskService {
	return &taskService{taskRepository: r}
}

func (s *taskService) AddTask(tasks models.Tasks) (models.Tasks, error) {
	if tasks.Title == "" || tasks.AssignedTo == "" {
		return models.Tasks{}, errors.New("title or AssignedTo не должны быть пустыми")
	}

	newTask, err := s.taskRepository.AddTask(tasks)
	if err != nil {
		return models.Tasks{}, err
	}

	return newTask, nil
}

func (s *taskService) ChangeTask(task models.Tasks) error {
	if task.ID == 0 {
		return errors.New("ID не должно быть пустым")
	}
	return s.taskRepository.ChangeTask(task)
}
