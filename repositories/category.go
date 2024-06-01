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

func (s *CategoryRepository) Get(user uint, id int) (models.Category, error) {
	var category models.Category

	error := s.DB.Preload("Recipes", "user_id = ?", user).First(&category, id).Error

	return category, error
}

func (s *CategoryRepository) GetAll(user uint) ([]models.Category, error) {
	var categories []models.Category

	error := s.DB.Joins("JOIN recipe_categories ON recipe_categories.category_id = categories.id").
		Joins("JOIN recipes ON recipes.id = recipe_categories.recipe_id").
		Where("recipes.user_id = ?", user).
		Group("categories.id").
		Find(&categories).Error

	return categories, error
}

func (s *CategoryRepository) Delete(id uint) error {
	return s.DB.Delete(&models.Category{}, id).Error
}
