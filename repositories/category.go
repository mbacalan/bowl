package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type CategoryRepository struct {
	db        *gorm.DB
	tableName string
}

func NewCategoryRepository(db *gorm.DB, tableName string) *CategoryRepository {
	repository := &CategoryRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s CategoryRepository) Get(id int) (models.Category, error) {
	var category models.Category

	error := s.db.Preload("Recipes").First(&category, id).Error

	return category, error
}

func (s CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	error := s.db.Preload("Recipes").First(&categories).Error

	return categories, error
}

func (s CategoryRepository) Delete(id uint) error {
	error := s.db.Delete(&models.Category{}, id).Error

	return error
}
