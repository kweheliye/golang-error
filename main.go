package main

import (
	"database/sql"
	"golang-error/internal/handler"
	"golang-error/internal/repository"
	"golang-error/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://admin:admin@localhost:5432/simplebank?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot open db connection: %v", err)
	}

	// Always test the actual connection:
	if err := db.Ping(); err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", (userHandler.GetByID))
	log.Print("Server Started")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
