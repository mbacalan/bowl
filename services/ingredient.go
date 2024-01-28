package services

import (
	"context"
	"log/slog"

	"github.com/mbacalan/bowl/db"
)

type IngredientService interface {
	Get(ctx context.Context) (ingredient db.Ingredient)
	GetAll(ctx context.Context) (ingredients []db.Ingredient)
	Create(ctx context.Context) (ingredient db.Ingredient)
}

type IngredientServiceImpl struct {
	Log             *slog.Logger
	IngredientStore *db.IngredientStore
}

func NewIngredientService(log *slog.Logger, rs *db.IngredientStore) IngredientServiceImpl {
	return IngredientServiceImpl{Log: log, IngredientStore: rs}
}

func (s *IngredientServiceImpl) Get(id int) (ingredient db.Ingredient, error error) {
	result, err := s.IngredientStore.GetIngredient(id)

	if err != nil {
		s.Log.Error("Error getting ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientServiceImpl) GetAll() (ingredients []db.Ingredient, error error) {
	result, err := s.IngredientStore.GetAllIngredients()

	if err != nil {
		s.Log.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *IngredientServiceImpl) Create(i db.Ingredient) (ingredient db.Ingredient, error error) {
	result, err := s.IngredientStore.CreateIngredient(i)

	if err != nil {
		s.Log.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}
