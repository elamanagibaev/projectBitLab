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

	userRepositories := repositories.NewUserRepositories(db)
	userService := services.NewUserService(userRepositories)
	userHandler := handlers.NewUserHandler(userService)

	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.AddUserHandler).Methods("POST")

	server := http.Server{
		Addr:    "localhost:4041",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
