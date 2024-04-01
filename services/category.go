package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
)

type CategoryService struct {
	Logger *slog.Logger
	Store  *repositories.CategoryRepository
}

func NewCategoryService(logger *slog.Logger, store *repositories.CategoryRepository) CategoryService {
	return CategoryService{
		Logger: logger,
		Store:  store,
	}
}

func (s *CategoryService) Get(id int) (repositories.Category, error) {
	result, err := s.Store.Get(id)

	if err != nil {
		s.Logger.Error("Error getting category", err)
		return result, err
	}

	return result, nil
}

func (s *CategoryService) GetAll() ([]repositories.Category, error) {
	result, err := s.Store.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all categories", err)
		return result, err
	}

	return result, nil
}
