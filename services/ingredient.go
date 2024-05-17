package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
)

type IngredientService struct {
	Logger     *slog.Logger
	Repository models.IngredientRepository
}

func NewIngredientService(logger *slog.Logger, repo models.IngredientRepository) *IngredientService {
	return &IngredientService{Logger: logger, Repository: repo}
}

func (s *IngredientService) Get(id int) (ingredient models.Ingredient, error error) {
	result, err := s.Repository.Get(id)

	if err != nil {
		s.Logger.Error("Error getting ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) GetAll() (ingredients []models.Ingredient, error error) {
	result, err := s.Repository.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) Create(i string) (ingredient models.Ingredient, error error) {
	result, err := s.Repository.GetOrCreate(i)

	if err != nil {
		s.Logger.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) Search(name string) (ingredients []models.Ingredient, error error) {
	if len(name) == 0 {
		return []models.Ingredient{}, error
	}

	result, err := s.Repository.Search(name)

	if err != nil {
		s.Logger.Error("Error searching ingredients", err)
		return result, err
	}

	return result, nil
}
