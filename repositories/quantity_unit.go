package repositories

import (
	"gorm.io/gorm"
)

type QuantityUnit struct {
	gorm.Model
	Name string
}

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

func (s QuantityUnitRepository) GetOrCreate(unit string) (QuantityUnit, error) {
	var entry QuantityUnit

	error := s.db.FirstOrCreate(&entry, QuantityUnit{Name: unit}).Error

	return entry, error
}

func (s QuantityUnitRepository) GetAll() ([]QuantityUnit, error) {
	var units []QuantityUnit

	error := s.db.Find(&units).Error

	return units, error
}
