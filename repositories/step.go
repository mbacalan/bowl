package repositories

import (
	"gorm.io/gorm"
)

type Step struct {
	gorm.Model
	RecipeID uint
	Step     string
}

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

func (s StepRepository) Create(step string, recipe uint) (i Step, err error) {
	entry := Step{Step: step, RecipeID: recipe}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s StepRepository) GetAll() (i []Step, err error) {
	var steps []Step

	error := s.db.Find(&steps).Error

	return steps, error
}

func (s StepRepository) Delete(id uint) (err error) {
	error := s.db.Delete(&Step{}, id).Error

	return error
}
