package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
)

type RecipeIngredientService struct {
	Log             *slog.Logger
	IngredientStore *db.RecipeIngredientRepository
}

func NewRecipeIngredientService(log *slog.Logger, rs *db.RecipeIngredientRepository) RecipeIngredientService {
	return RecipeIngredientService{Log: log, IngredientStore: rs}
}

func (s *RecipeIngredientService) GetAll() (ingredients []db.RecipeIngredient, error error) {
	result, err := s.IngredientStore.GetAll()

	if err != nil {
		s.Log.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeIngredientService) Create(r uint, i uint, qu uint, q string) (ingredient db.RecipeIngredient, error error) {
	result, err := s.IngredientStore.Create(r, i, qu, q)

	if err != nil {
		s.Log.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}
