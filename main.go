package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/mbacalan/bowl/handlers"
	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services"
	"gorm.io/driver/sqlite"
)

func main() {
	godotenv.Load()
	server := createServer()
	handlers.MountHandlers(server)

	fmt.Printf("Listening on %v\n", ":3000")

	if err := http.ListenAndServe(":3000", server.Router); err != nil {
		server.Logger.Error("Failed to start the server", err)
		os.Exit(1)
	}
}

func createServer() *models.Server {
	db, err := repositories.NewConnection(sqlite.Open("./db.sqlite"))
	repositories.SeedQuantityUnits(db)
	if err != nil {
		os.Exit(1)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repos := repositories.CreateRepositories(db)
	services := services.CreateServices(logger, repos)
	handlers := handlers.CreateHandlers(logger, services)

	return &models.Server{
		Database: db,
		Router:   chi.NewRouter(),
		Logger:   slog.New(slog.NewJSONHandler(os.Stdout, nil)),

		Repos:    repos,
		Services: services,
		Handlers: handlers,
	}
}
