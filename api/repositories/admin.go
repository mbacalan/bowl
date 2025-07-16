package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type AdminRepository struct {
	*models.Repository
}

func NewAdminRepository(db *gorm.DB, tableName string) *AdminRepository {
	return &AdminRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *AdminRepository) Get(user uint) (models.User, error) {
	var entry models.User

	error := s.DB.First(&entry, user).Error

	return entry, error
}

func (s *AdminRepository) GetIngredients() ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	error := s.DB.Find(&ingredients).Error

	return ingredients, error
}

func (s *AdminRepository) CreateIngredient(ingredient string) (models.Ingredient, error) {
	var entry models.Ingredient

	error := s.DB.FirstOrCreate(&entry, models.Ingredient{Name: ingredient}).Error

	return entry, error
}

func (s *AdminRepository) DeleteIngredient(id uint) (bool, error) {
	var entry models.Ingredient

	error := s.DB.Where("id = ?", id).Delete(&entry).Error

	return error == nil, error
}

func (s *AdminRepository) GetQuantityUnits() ([]models.QuantityUnit, error) {
	var units []models.QuantityUnit

	error := s.DB.Find(&units).Error

	return units, error
}

func (s *AdminRepository) CreateQuantityUnit(unit string) (models.QuantityUnit, error) {
	var entry models.QuantityUnit

	error := s.DB.FirstOrCreate(&entry, models.QuantityUnit{Name: unit}).Error

	return entry, error
}

func (s *AdminRepository) DeleteQuantityUnit(id uint) (bool, error) {
	var entry models.QuantityUnit

	error := s.DB.Where("id = ?", id).Delete(&entry).Error

	return error == nil, error
}
