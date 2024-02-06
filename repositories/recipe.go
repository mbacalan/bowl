package db

import (
	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	Name              string
	RecipeIngredients []RecipeIngredient
	// Steps       []string
	// Rating      uint
	// Difficulty  uint
	// Duration    uint
	// Categories  []Category
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

func (s RecipeRepository) CreateRecipe(recipe Recipe) (r Recipe, err error) {
	entry := Recipe{Name: recipe.Name}
	result := s.DB.Create(&entry)

	if result.Error != nil {
		return Recipe{}, result.Error
	}

	return entry, nil
}

func (s RecipeRepository) GetRecipe(id int) (r Recipe, err error) {
	var recipe Recipe
	result := s.DB.Model(&Recipe{}).Preload("RecipeIngredients").Preload("RecipeIngredients.Ingredient").Preload("RecipeIngredients.QuantityUnit").Find(&recipe, id)

	if result.Error != nil {
		return Recipe{}, result.Error
	}

	return recipe, nil
}

func (s RecipeRepository) GetAllRecipes() (r []Recipe, err error) {
	var recipes []Recipe
	result := s.DB.Find(&recipes)

	if result.Error != nil {
		return []Recipe{}, result.Error
	}

	return recipes, nil
}

func (s RecipeRepository) GetRecentRecipes(limit int) (r []Recipe, err error) {
	var recipes []Recipe
	result := s.DB.Order("id DESC").Limit(limit).Find(&recipes)

	if result.Error != nil {
		return []Recipe{}, result.Error
	}

	return recipes, nil
}
