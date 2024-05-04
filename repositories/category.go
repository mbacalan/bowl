package repositories

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name    string
	Recipes []*Recipe `gorm:"many2many:recipe_categories;"`
}

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

func (s CategoryRepository) Get(id int) (Category, error) {
	var category Category

	error := s.db.Preload("Recipes").First(&category, id).Error

	return category, error
}

func (s CategoryRepository) GetAll() ([]Category, error) {
	var categories []Category

	error := s.db.Preload("Recipes").First(&categories).Error

	return categories, error
}

func (s CategoryRepository) Delete(id uint) error {
	error := s.db.Delete(&Category{}, id).Error

	return error
}
