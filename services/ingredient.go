package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
)

type IngredientService struct {
	Log             *slog.Logger
	IngredientStore *db.IngredientRepository
}

func NewIngredientService(log *slog.Logger, rs *db.IngredientRepository) IngredientService {
	return IngredientService{Log: log, IngredientStore: rs}
}

func (s *IngredientService) Get(id int) (ingredient db.Ingredient, error error) {
	result, err := s.IngredientStore.GetIngredient(id)

	if err != nil {
		s.Log.Error("Error getting ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) GetAll() (ingredients []db.Ingredient, error error) {
	result, err := s.IngredientStore.GetAllIngredients()

	if err != nil {
		s.Log.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) Create(i string) (ingredient db.Ingredient, error error) {
	result, err := s.IngredientStore.GetOrCreate(i)

	if err != nil {
		s.Log.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientService) Search(name string) (ingredients []db.Ingredient, error error) {
	if len(name) == 0 {
		return []db.Ingredient{}, error
	}

	result, err := s.IngredientStore.SearchIngredient(name)

	if err != nil {
		s.Log.Error("Error searching ingredients", err)
		return result, err
	}

	return result, nil
}
