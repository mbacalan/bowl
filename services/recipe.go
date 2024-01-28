package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/db"
)

type RecipeService struct {
	Log         *slog.Logger
	RecipeStore *db.RecipeStore
}

func NewRecipeService(log *slog.Logger, rs *db.RecipeStore) RecipeService {
	return RecipeService{Log: log, RecipeStore: rs}
}

func (s *RecipeService) Get(id int) (recipe db.Recipe, error error) {
	result, err := s.RecipeStore.GetRecipe(id)

	if err != nil {
		s.Log.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetAll() (recipes []db.Recipe, error error) {
	result, err := s.RecipeStore.GetAllRecipes()

	if err != nil {
		s.Log.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetRecent(limit int) (recipes []db.Recipe, error error) {
	result, err := s.RecipeStore.GetRecentRecipes(limit)

	if err != nil {
		s.Log.Error("Error getting recent recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) Create(r db.Recipe) (recipe db.Recipe, error error) {
	result, err := s.RecipeStore.CreateRecipe(r)

	if err != nil {
		s.Log.Error("Error creating recipe", err)
		return result, err
	}

	return result, nil
}
