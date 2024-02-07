package db

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
	result := s.db.Create(&entry)

	if result.Error != nil {
		return Step{}, result.Error
	}

	return entry, nil
}

func (s StepRepository) GetAll() (i []Step, err error) {
	var steps []Step
	result := s.db.Find(&steps)

	if result.Error != nil {
		return []Step{}, result.Error
	}

	return steps, nil
}
