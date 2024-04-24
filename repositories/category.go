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
	result := s.db.Preload("Recipes").Find(&category, id)

	if result.Error != nil {
		return Category{}, result.Error
	}

	return category, nil
}

func (s CategoryRepository) GetOrCreate(category string, recipe uint) (Category, error) {
	var entry Category
	result := s.db.FirstOrCreate(&entry, Category{Name: category})

	if result.Error != nil {
		return Category{}, result.Error
	}

	return entry, nil
}

func (s CategoryRepository) GetAll() ([]Category, error) {
	var categories []Category

	error := s.db.Preload("Recipes").Find(&categories).Error

	return categories, error
}

func (s CategoryRepository) Delete(id uint) error {
	result := s.db.Delete(&Category{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
