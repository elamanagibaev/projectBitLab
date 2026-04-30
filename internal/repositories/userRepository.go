package repositories

import (
	"database/sql"
	"projectBit/models"
)

type UserRepository interface {
	AddUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) AddUser(user models.User) (models.User, error) {
	err := r.db.QueryRow("INSERT INTO users(name, email) VALUES ($1, $2) RETURNING id, created_at",
		user.Name,
		user.Email,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
