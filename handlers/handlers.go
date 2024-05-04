package handlers

import (
	"log/slog"

	"github.com/mbacalan/bowl/services"
)

type Handlers struct {
	HomeHandler         *HomeHandler
	RecipeHandler       *RecipeHandler
	IngredientHandler   *IngredientHandler
	QuantityUnitHandler *QuantityUnitHandler
	CategoryHandler     *CategoryHandler
}

func CreateHandlers(logger *slog.Logger, services *services.Services) *Handlers {
	return &Handlers{
		HomeHandler:         NewHomeHandler(logger, services.RecipeService),
		RecipeHandler:       NewRecipeHandler(logger, services.RecipeService),
		IngredientHandler:   NewIngredientHandler(logger, services.IngredientService),
		QuantityUnitHandler: NewQuantityUnitHandler(logger, services.QuantityUnitService),
		CategoryHandler:     NewCategoryHandler(logger, services.CategoryService),
	}
}
