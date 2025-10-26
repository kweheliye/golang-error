package main

import (
	"database/sql"
	"golang-error/internal/handler"
	"golang-error/internal/service"
	"golang-error/internal/store"
	"golang-error/util"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := util.LoadConfig("../")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Printf("config: %v", cfg)

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	)

	db, err := sql.Open("postgres", "postgresql://admin:admin@localhost:5432/simplebank?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot open db connection: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}
	
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("cannot close db connection: %v", err)
		}
	}(db)

	// --- Dependency Injection --- //
	sqlStore := store.NewSQLStore(db)
	services := service.NewService(sqlStore)
	// The handler now sets up all the routes.
	h := handler.NewHandler(services, logger)

	log.Printf("Starting server on:%s", cfg.Server.Port)

	// We pass the handler's configured router to the server.
	if err := http.ListenAndServe(cfg.Server.Port, h.Router()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
