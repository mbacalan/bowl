package db

import (
	"gorm.io/gorm"
)

type QuantityUnit struct {
	gorm.Model
	Unit string
}

type QuantityUnitStore struct {
	db        *gorm.DB
	tableName string
}

func NewQuantityUnitStore(db *gorm.DB, tableName string) *QuantityUnitStore {
	store := &QuantityUnitStore{
		tableName: tableName,
		db:        db,
	}

	return store
}

func (s QuantityUnitStore) GetOrCreate(unit string) (i QuantityUnit, err error) {
	var entry QuantityUnit
	result := s.db.FirstOrCreate(&entry, QuantityUnit{Unit: unit})

	if result.Error != nil {
		return QuantityUnit{}, result.Error
	}

	return entry, nil
}

func (s QuantityUnitStore) GetAll() (i []QuantityUnit, err error) {
	var units []QuantityUnit
	result := s.db.Find(&units)

	if result.Error != nil {
		return []QuantityUnit{}, result.Error
	}

	return units, nil
}
