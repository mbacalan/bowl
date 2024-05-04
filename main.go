package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mbacalan/bowl/assets"
	"github.com/mbacalan/bowl/handlers"
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
	Handlers *Handlers
}

type Handlers struct {
	HomeHandler         *handlers.HomeHandler
	RecipeHandler       *handlers.RecipeHandler
	IngredientHandler   *handlers.IngredientHandler
	QuantityUnitHandler *handlers.QuantityUnitHandler
	CategoryHandler     *handlers.CategoryHandler
}

func main() {
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
	s.createHandlers()

	return s
}

func (s *Server) createHandlers() {
	s.Handlers = &Handlers{
		HomeHandler:         handlers.NewHomeHandler(s.Logger, s.Services.RecipeService),
		RecipeHandler:       handlers.NewRecipeHandler(s.Logger, s.Services.RecipeService),
		IngredientHandler:   handlers.NewIngredientHandler(s.Logger, s.Services.IngredientService),
		QuantityUnitHandler: handlers.NewQuantityUnitHandler(s.Logger, s.Services.QuantityUnitService),
		CategoryHandler:     handlers.NewCategoryHandler(s.Logger, s.Services.CategoryService),
	}
}

func (s *Server) mountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Compress(5))

	s.Router.Mount("/assets", assets.Routes())

	s.Router.Mount("/", s.Handlers.HomeHandler.Routes())
	s.Router.Mount("/recipes", s.Handlers.RecipeHandler.Routes())
	s.Router.Mount("/categories", s.Handlers.CategoryHandler.Routes())
	s.Router.Mount("/ingredients", s.Handlers.IngredientHandler.Routes())
	s.Router.Mount("/quantity-units", s.Handlers.QuantityUnitHandler.Routes())
}
