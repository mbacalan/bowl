package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type QuantityUnitRepository struct {
	*models.Repository
}

func NewQuantityUnitRepository(db *gorm.DB, tableName string) *QuantityUnitRepository {
	return &QuantityUnitRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *QuantityUnitRepository) GetOrCreate(unit string) (models.QuantityUnit, error) {
	var entry models.QuantityUnit

	error := s.DB.FirstOrCreate(&entry, models.QuantityUnit{Name: unit}).Error

	return entry, error
}

func (s *QuantityUnitRepository) GetAll() ([]models.QuantityUnit, error) {
	var units []models.QuantityUnit

	error := s.DB.Find(&units).Error

	return units, error
}
