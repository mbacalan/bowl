package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

type StepRepository struct {
	db        *gorm.DB
	tableName string
}

func NewStepRepository(db *gorm.DB, tableName string) *StepRepository {
	repository := &StepRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s StepRepository) Create(step string, recipe uint) (models.Step, error) {
	entry := models.Step{Step: step, RecipeID: recipe}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s StepRepository) GetAll() ([]models.Step, error) {
	var steps []models.Step

	error := s.db.Find(&steps).Error

	return steps, error
}

func (s StepRepository) Delete(id uint) error {
	error := s.db.Delete(&models.Step{}, id).Error

	return error
}
