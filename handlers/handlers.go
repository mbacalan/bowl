package handlers

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mbacalan/bowl/assets"
	"github.com/mbacalan/bowl/handlers/internal"
	"github.com/mbacalan/bowl/models"
)

func CreateHandlers(logger *slog.Logger, services *models.Services) *models.Handlers {
	authHandler := NewAuthHandler(logger, services.AuthService)
	store := authHandler.GetStore()

	return &models.Handlers{
		AuthHandler:         authHandler,
		HomeHandler:         NewHomeHandler(logger, services.RecipeService, store),
		RecipeHandler:       NewRecipeHandler(logger, services.RecipeService, store),
		IngredientHandler:   NewIngredientHandler(logger, services.IngredientService),
		QuantityUnitHandler: NewQuantityUnitHandler(logger, services.QuantityUnitService),
		CategoryHandler:     NewCategoryHandler(logger, services.CategoryService, store),
	}
}

func MountHandlers(s *models.Server) {
	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))

		r.Mount("/assets", assets.Routes())
		r.Mount("/auth", s.Handlers.AuthHandler.Routes())
	})

	s.Router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(middleware.Compress(5))
		r.Use(internal.Authenticated(s.Handlers.AuthHandler.GetStore()))

		r.Mount("/", s.Handlers.HomeHandler.Routes())
		r.Mount("/recipes", s.Handlers.RecipeHandler.Routes())
		r.Mount("/categories", s.Handlers.CategoryHandler.Routes())
		r.Mount("/ingredients", s.Handlers.IngredientHandler.Routes())
		r.Mount("/quantity-units", s.Handlers.QuantityUnitHandler.Routes())
	})
}
