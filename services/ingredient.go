package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/db"
)

type IngredientService struct {
	Log             *slog.Logger
	IngredientStore *db.IngredientStore
}

func NewIngredientService(log *slog.Logger, rs *db.IngredientStore) IngredientService {
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

func (s *IngredientService) Create(i db.Ingredient) (ingredient db.Ingredient, error error) {
	result, err := s.IngredientStore.CreateIngredient(i)

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
