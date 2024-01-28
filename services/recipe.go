package services

import (
	"context"
	"log/slog"

	"github.com/mbacalan/bowl/db"
)

type RecipeService interface {
	Get(ctx context.Context) (recipe db.Recipe)
	GetAll(ctx context.Context) (recipes []db.Recipe)
}

type RecipeServiceImpl struct {
	Log         *slog.Logger
	RecipeStore *db.RecipeStore
}

func NewRecipeService(log *slog.Logger, rs *db.RecipeStore) RecipeServiceImpl {
	return RecipeServiceImpl{Log: log, RecipeStore: rs}
}

func (s *RecipeServiceImpl) Get(id int) (recipe db.Recipe, error error) {
	result, err := s.RecipeStore.GetRecipe(id)

	if err != nil {
		s.Log.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeServiceImpl) GetAll() (recipes []db.Recipe, error error) {
	result, err := s.RecipeStore.GetAllRecipes()

	if err != nil {
		s.Log.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeServiceImpl) GetRecent(limit int) (recipes []db.Recipe, error error) {
	result, err := s.RecipeStore.GetRecentRecipes(limit)

	if err != nil {
		s.Log.Error("Error getting recent recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeServiceImpl) Create(r db.Recipe) (recipe db.Recipe, error error) {
	result, err := s.RecipeStore.CreateRecipe(r)

	if err != nil {
		s.Log.Error("Error creating recipe", err)
		return result, err
	}

	return result, nil
}
