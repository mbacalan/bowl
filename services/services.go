package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
)

type Services struct {
	RecipeService       *RecipeService
	IngredientService   *IngredientService
	QuantityUnitService *QuantityUnitService
	CategoryService     *CategoryService
}

func CreateServices(logger *slog.Logger, repos *repositories.Repositories) *Services {
	return &Services{
		RecipeService:       NewRecipeService(logger, repos.RecipeRepository),
		IngredientService:   NewIngredientService(logger, repos.IngredientRepository),
		QuantityUnitService: NewQuantityUnitService(logger, repos.QuantityUnitRepository),
		CategoryService:     NewCategoryService(logger, repos.CategoryRepository),
	}
}
