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

func (s QuantityUnitRepository) GetOrCreate(unit string) (i QuantityUnit, err error) {
	var entry QuantityUnit
	result := s.db.FirstOrCreate(&entry, QuantityUnit{Name: unit})

	if result.Error != nil {
		return QuantityUnit{}, result.Error
	}

	return entry, nil
}

func (s QuantityUnitRepository) GetAll() (i []QuantityUnit, err error) {
	var units []QuantityUnit
	result := s.db.Find(&units)

	if result.Error != nil {
		return []QuantityUnit{}, result.Error
	}

	return units, nil
}
