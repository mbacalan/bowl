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

type Service struct {
	Log         *slog.Logger
	RecipeStore *db.RecipeStore
}

func New(log *slog.Logger, rs *db.RecipeStore) Service {
	return Service{Log: log, RecipeStore: rs}
}

func (s *Service) Get(id int) (recipe db.Recipe, error error) {
	result, err := s.RecipeStore.Get(id)

	if err != nil {
		s.Log.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *Service) GetAll() (recipes []db.Recipe, error error) {
	result, err := s.RecipeStore.GetAll()

	if err != nil {
		s.Log.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}
