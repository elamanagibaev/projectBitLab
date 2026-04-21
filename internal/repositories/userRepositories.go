package repositories

import (
	"database/sql"
	"log"
	"projectBit/models"
)

type UserRepositories interface {
	AddUser(user models.User)
}

type userRepositories struct {
	db *sql.DB
}

func NewUserRepositories(db *sql.DB) UserRepositories {
	return &userRepositories{db: db}
}

func (userRepositories *userRepositories) AddUser(user models.User) {
	_, err := userRepositories.db.Exec("insert into users(name, email, created_at) values ($1, $2, $3)", user.Name, user.Email, user.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
}
