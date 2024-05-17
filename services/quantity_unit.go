package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
)

type QuantityUnitService struct {
	Logger     *slog.Logger
	Repository models.QuantityUnitRepository
}

func NewQuantityUnitService(logger *slog.Logger, repo models.QuantityUnitRepository) *QuantityUnitService {
	return &QuantityUnitService{Logger: logger, Repository: repo}
}

func (s *QuantityUnitService) GetAll() (units []models.QuantityUnit, error error) {
	result, err := s.Repository.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all quantity units", err)
		return result, err
	}

	return result, nil
}

func (s *QuantityUnitService) Create(i string) (unit models.QuantityUnit, error error) {
	result, err := s.Repository.GetOrCreate(i)

	if err != nil {
		s.Logger.Error("Error creating unit", err)
		return result, err
	}

	return result, nil
}
