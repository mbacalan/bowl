package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

type QuantityUnitRepository struct {
	db        *gorm.DB
	tableName string
}

func NewQuantityUnitRepository(db *gorm.DB, tableName string) *QuantityUnitRepository {
	repository := &QuantityUnitRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s QuantityUnitRepository) GetOrCreate(unit string) (models.QuantityUnit, error) {
	var entry models.QuantityUnit

	error := s.db.FirstOrCreate(&entry, models.QuantityUnit{Name: unit}).Error

	return entry, error
}

func (s QuantityUnitRepository) GetAll() ([]models.QuantityUnit, error) {
	var units []models.QuantityUnit

	error := s.db.Find(&units).Error

	return units, error
}
