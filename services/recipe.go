package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
)

type RecipeService struct {
	Log                        *slog.Logger
	RecipeRepository           *db.RecipeRepository
	RecipeIngredientRepository *db.RecipeIngredientRepository
	IngredientRepository       *db.IngredientRepository
	QuantityUnitRepository     *db.QuantityUnitRepository
}

func NewRecipeService(log *slog.Logger, rs *db.RecipeRepository, rids *db.RecipeIngredientRepository, ids *db.IngredientRepository, qds *db.QuantityUnitRepository) RecipeService {
	return RecipeService{
		Log:                        log,
		RecipeRepository:           rs,
		RecipeIngredientRepository: rids,
		IngredientRepository:       ids,
		QuantityUnitRepository:     qds,
	}
}

func (s *RecipeService) Get(id int) (recipe db.Recipe, error error) {
	result, err := s.RecipeRepository.GetRecipe(id)

	if err != nil {
		s.Log.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetAll() (recipes []db.Recipe, error error) {
	result, err := s.RecipeRepository.GetAllRecipes()

	if err != nil {
		s.Log.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetRecent(limit int) (recipes []db.Recipe, error error) {
	result, err := s.RecipeRepository.GetRecentRecipes(limit)

	if err != nil {
		s.Log.Error("Error getting recent recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) Create(r db.Recipe) (recipe db.Recipe, error error) {
	result, err := s.RecipeRepository.CreateRecipe(r)

	if err != nil {
		s.Log.Error("Error creating recipe", err)
		return result, err
	}

	return result, nil
}
