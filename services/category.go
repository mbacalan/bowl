package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
)

type CategoryService struct {
	Logger     *slog.Logger
	Repository models.CategoryRepository
}

func NewCategoryService(logger *slog.Logger, repo models.CategoryRepository) *CategoryService {
	return &CategoryService{
		Logger:     logger,
		Repository: repo,
	}
}

func (s *CategoryService) Get(id int) (models.Category, error) {
	result, err := s.Repository.Get(id)

	if err != nil {
		s.Logger.Error("Error getting category", err)
		return result, err
	}

	return result, nil
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	result, err := s.Repository.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all categories", err)
		return result, err
	}

	return result, nil
}
