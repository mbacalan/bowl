package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
)

type AdminService struct {
	Logger     *slog.Logger
	Repository models.AdminRepository
}

func NewAdminService(logger *slog.Logger, repo models.AdminRepository) *AdminService {
	return &AdminService{
		Logger:     logger,
		Repository: repo,
	}
}

func (s *AdminService) IsAdmin(user uint) bool {
	dbuser, _ := s.Repository.Get(user)

	if dbuser.IsAdmin {
		return true
	} else {
		return false
	}
}

func (s *AdminService) GetIngredients() ([]models.Ingredient, error) {
	result, err := s.Repository.GetIngredients()

	if err != nil {
		s.Logger.Error("Error getting all ingredients", err)
		return result, err
	}

	return result, nil
}

func (s *AdminService) CreateIngredient(name string) (models.Ingredient, error) {
	result, err := s.Repository.CreateIngredient(name)

	if err != nil {
		s.Logger.Error("Error creating ingredient", err)
		return result, err
	}

	return result, nil
}

func (s *AdminService) GetQuantityUnits() ([]models.QuantityUnit, error) {
	result, err := s.Repository.GetQuantityUnits()

	if err != nil {
		s.Logger.Error("Error getting all quantity units", err)
		return result, err
	}

	return result, nil
}

func (s *AdminService) CreateQuantityUnit(name string) (models.QuantityUnit, error) {
	result, err := s.Repository.CreateQuantityUnit(name)

	if err != nil {
		s.Logger.Error("Error creating quantity unit", err)
		return result, err
	}

	return result, nil
}
