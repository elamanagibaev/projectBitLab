package main

import (
	"database/sql"
	"log"
	"net/http"
	"projectBit/internal/handlers"
	"projectBit/internal/repositories"
	"projectBit/internal/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	connection := "user=postgres password=Elaman2004123 dbname=BitLabProject host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}

	errConnect := db.Ping()
	if errConnect != nil {
		log.Fatal(errConnect)
	}
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	InitDB()
	defer CloseDB()
	router := mux.NewRouter()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.AddUser).Methods("POST")

	projectRepository := repositories.NewProjectRepository(db)
	projectService := services.NewProjectService(projectRepository)
	projectHandler := handlers.NewProjectHandler(projectService)

	router.HandleFunc("/projects", projectHandler.AddProject).Methods("POST")

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	router.HandleFunc("/task", taskHandler.AddTask).Methods("POST")
	router.HandleFunc("/task", taskHandler.ChangeTask).Methods("PUT")

	server := http.Server{
		Addr:    "localhost:4041",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
