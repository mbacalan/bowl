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

func (s *CategoryService) Get(user uint, id int) (models.Category, error) {
	result, err := s.Repository.Get(user, id)

	if err != nil {
		s.Logger.Error("Error getting category", "error", err)
		return result, err
	}

	return result, nil
}

func (s *CategoryService) GetAll(user uint) ([]models.Category, error) {
	result, err := s.Repository.GetAll(user)

	if err != nil {
		s.Logger.Error("Error getting all categories", "error", err)
		return result, err
	}

	return result, nil
}
