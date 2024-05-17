package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type CategoryRepository struct {
	*models.Repository
}

func NewCategoryRepository(db *gorm.DB, tableName string) *CategoryRepository {
	return &CategoryRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *CategoryRepository) Get(id int) (models.Category, error) {
	var category models.Category

	error := s.DB.Preload("Recipes").First(&category, id).Error

	return category, error
}

func (s *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category

	error := s.DB.Preload("Recipes").Find(&categories).Error

	return categories, error
}

func (s *CategoryRepository) Delete(id uint) error {
	return s.DB.Delete(&models.Category{}, id).Error
}
