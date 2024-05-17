package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/mbacalan/bowl/assets"
	"github.com/mbacalan/bowl/handlers"
	"github.com/mbacalan/bowl/internal"
	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	Router   *chi.Mux
	Database *gorm.DB
	Logger   *slog.Logger
	Repos    *repositories.Repositories
	Services *services.Services
	Handlers *handlers.Handlers
}

func main() {
	godotenv.Load()
	server := createServer()
	server.mountHandlers()

	fmt.Printf("Listening on %v\n", ":3000")

	if err := http.ListenAndServe(":3000", server.Router); err != nil {
		server.Logger.Error("Failed to start the server", err)
		os.Exit(1)
	}
}

func createServer() *Server {
	s := &Server{}
	db, err := repositories.NewConnection(sqlite.Open("./db.sqlite"))
	repositories.SeedQuantityUnits(db)
	if err != nil {
		os.Exit(1)
	}

	s.Database = db
	s.Router = chi.NewRouter()
	s.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s.Repos = repositories.CreateRepositories(s.Database)
	s.Services = services.CreateServices(s.Logger, s.Repos)
	s.Handlers = handlers.CreateHandlers(s.Logger, s.Services)

	return s
}

func (s *Server) mountHandlers() {
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))

		r.Mount("/assets", assets.Routes())
		r.Mount("/auth", s.Handlers.AuthHandler.Routes())
	})

	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))
		r.Use(internal.Authenticated(s.Handlers.AuthHandler.Store))

		r.Mount("/", s.Handlers.HomeHandler.Routes())
		r.Mount("/recipes", s.Handlers.RecipeHandler.Routes())
		r.Mount("/categories", s.Handlers.CategoryHandler.Routes())
		r.Mount("/ingredients", s.Handlers.IngredientHandler.Routes())
		r.Mount("/quantity-units", s.Handlers.QuantityUnitHandler.Routes())
	})
}
