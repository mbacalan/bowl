package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
)

type RecipeIngredientService struct {
	Logger     *slog.Logger
	Repository models.RecipeIngredientRepository
}

func NewRecipeIngredientService(logger *slog.Logger, repo models.RecipeIngredientRepository) *RecipeIngredientService {
	return &RecipeIngredientService{Logger: logger, Repository: repo}
}

func (s *RecipeIngredientService) GetAll() (ingredients []models.RecipeIngredient, error error) {
	result, err := s.Repository.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeIngredientService) Create(r uint, i uint, qu uint, q string) (ingredient models.RecipeIngredient, error error) {
	result, err := s.Repository.Create(r, i, qu, q)

	if err != nil {
		s.Logger.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}
