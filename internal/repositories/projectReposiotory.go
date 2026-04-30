package repositories

import (
	"database/sql"
	"projectBit/models"
)

type ProjectRepository interface {
	AddProject(project models.Projects) (models.Projects, error)
}

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) AddProject(project models.Projects) (models.Projects, error) {
	err := r.db.QueryRow("INSERT INTO projects(title, description) VALUES ($1, $2) RETURNING id, created_at", project.Title, project.Description).Scan(&project.ID, &project.CreatedAt)

	if err != nil {
		return models.Projects{}, err
	}

	return project, nil
}
