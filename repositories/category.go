package db

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	RecipeID uint
	Category string
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

func (s CategoryRepository) Create(category string, recipe uint) (i Category, err error) {
	entry := Category{Category: category, RecipeID: recipe}
	result := s.db.Create(&entry)

	if result.Error != nil {
		return Category{}, result.Error
	}

	return entry, nil
}

func (s CategoryRepository) GetAll() (i []Category, err error) {
	var categories []Category
	result := s.db.Find(&categories)

	if result.Error != nil {
		return []Category{}, result.Error
	}

	return categories, nil
}

func (s CategoryRepository) Delete(id uint) (err error) {
	result := s.db.Delete(&Category{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
