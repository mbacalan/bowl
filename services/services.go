package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services/internal"
)

type Services struct {
	AuthService         *AuthService
	RecipeService       *RecipeService
	IngredientService   *IngredientService
	QuantityUnitService *QuantityUnitService
	CategoryService     *CategoryService
}

func CreateServices(logger *slog.Logger, repos *repositories.Repositories) *Services {
	hash := internal.NewArgon2idHash(1, 32, 64*1024, 32, 256)

	return &Services{
		AuthService:         NewAuthService(logger, repos.UserRepository, hash),
		RecipeService:       NewRecipeService(logger, repos.RecipeRepository),
		IngredientService:   NewIngredientService(logger, repos.IngredientRepository),
		QuantityUnitService: NewQuantityUnitService(logger, repos.QuantityUnitRepository),
		CategoryService:     NewCategoryService(logger, repos.CategoryRepository),
	}
}
