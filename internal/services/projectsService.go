package services

import (
	"errors"
	"projectBit/internal/repositories"
	"projectBit/models"
)

type ProjectService interface {
	AddProject(project models.Projects) (models.Projects, error)
}

type projectService struct {
	projectRepository repositories.ProjectRepository
}

func NewProjectService(projectRepository repositories.ProjectRepository) ProjectService {
	return &projectService{projectRepository: projectRepository}
}

func (s *projectService) AddProject(project models.Projects) (models.Projects, error) {
	if project.Title == "" || project.Description == "" {
		return models.Projects{}, errors.New("title и description обязательны")
	}
	newProject, err := s.projectRepository.AddProject(project)

	if err != nil {
		return models.Projects{}, err
	}
	return newProject, nil
}
