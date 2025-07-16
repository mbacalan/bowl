package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type StepRepository struct {
	*models.Repository
}

func NewStepRepository(db *gorm.DB, tableName string) *StepRepository {
	return &StepRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *StepRepository) Create(step string, recipe uint) (models.Step, error) {
	entry := models.Step{Step: step, RecipeID: recipe}

	error := s.DB.Create(&entry).Error

	return entry, error
}

func (s *StepRepository) GetAll() ([]models.Step, error) {
	var steps []models.Step

	error := s.DB.Find(&steps).Error

	return steps, error
}

func (s *StepRepository) Delete(id uint) error {
	error := s.DB.Delete(&models.Step{}, id).Error

	return error
}
