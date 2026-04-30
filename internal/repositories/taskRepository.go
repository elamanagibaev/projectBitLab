package repositories

import (
	"database/sql"
	"projectBit/models"
)

type TaskRepository interface {
	AddTask(task models.Tasks) (models.Tasks, error)
	ChangeTask(task models.Tasks) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) AddTask(task models.Tasks) (models.Tasks, error) {
	err := r.db.QueryRow("INSERT INTO tasks(title, description, project_id, assigned_to) VALUES ($1, $2, $3, $4) RETURNING id, status, created_at",
		task.Title, task.Description, task.ProjectId, task.AssignedTo).Scan(&task.ID, &task.Status, &task.CreatedAt)

	if err != nil {
		return models.Tasks{}, err
	}

	return task, nil
}

func (r *taskRepository) ChangeTask(task models.Tasks) error {
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, project_id = $4, assigned_to = $5 WHERE id = $6",
		task.Title, task.Description, task.Status, task.ProjectId, task.AssignedTo, task.ID)
	return err
}
