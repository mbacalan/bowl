package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/services/internal"
)

func CreateServices(logger *slog.Logger, repos *models.Repositories) *models.Services {
	hash := internal.NewArgon2idHash(1, 32, 64*1024, 32, 256)

	return &models.Services{
		AuthService:         NewAuthService(logger, repos.UserRepository, hash),
		AdminService:        NewAdminService(logger, repos.AdminRepository),
		RecipeService:       NewRecipeService(logger, repos.RecipeRepository),
		QuantityUnitService: NewQuantityUnitService(logger, repos.QuantityUnitRepository),
		CategoryService:     NewCategoryService(logger, repos.CategoryRepository),
	}
}
