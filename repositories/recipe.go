package repositories

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name              string
	RecipeIngredients []RecipeIngredient
	Steps             []Step
	PrepDuration      uint
	CookDuration      uint
	Categories        []*Category `gorm:"many2many:recipe_categories;"`
	// Rating      uint
	// Difficulty  uint
	// WwPoints    uint `gorm:"column:ww_points"`
	// Language    string
	// Comments    []string
}

type RecipeRepository struct {
	DB        *gorm.DB
	tableName string
}

func NewRecipeRepository(db *gorm.DB, tableName string) *RecipeRepository {
	repository := &RecipeRepository{
		tableName: tableName,
		DB:        db,
	}

	return repository
}

func (s RecipeRepository) Create(name string, prep uint, cook uint) (Recipe, error) {
	entry := Recipe{
		Name:         name,
		PrepDuration: prep,
		CookDuration: cook,
	}

	error := s.DB.Create(&entry).Error

	return entry, error
}

func (s RecipeRepository) Get(id int) (Recipe, error) {
	var recipe Recipe

	error := s.DB.Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.QuantityUnit").
		Preload("Steps").
		Preload("Categories").
		First(&recipe, id).Error

	return recipe, error
}

func (s RecipeRepository) GetAll() (r []Recipe, err error) {
	var recipes []Recipe

	error := s.DB.Find(&recipes).Error

	return recipes, error
}

func (s RecipeRepository) GetRecent(limit int) (r []Recipe, err error) {
	var recipes []Recipe

	error := s.DB.Order("id DESC").Limit(limit).Find(&recipes).Error

	return recipes, error
}
