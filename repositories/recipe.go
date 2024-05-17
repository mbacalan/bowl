package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type RecipeRepository struct {
	*models.Repository
}

func NewRecipeRepository(db *gorm.DB, tableName string) *RecipeRepository {
	return &RecipeRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func NewRecipeUOW(database *gorm.DB) models.RecipeUnitOfWork {
	return models.RecipeUnitOfWork{
		DB:                         database,
		RecipeRepository:           NewRecipeRepository(database, "recipes"),
		IngredientRepository:       NewIngredientRepository(database, "ingredients"),
		QuantityUnitRepository:     NewQuantityUnitRepository(database, "quantity_units"),
		RecipeIngredientRepository: NewRecipeIngredientRepository(database, "recipe_ingredients"),
		StepRepository:             NewStepRepository(database, "steps"),
		CategoryRepository:         NewCategoryRepository(database, "categories"),
	}
}

func (s *RecipeRepository) Create(name string, prep uint, cook uint) (models.Recipe, error) {
	entry := models.Recipe{
		Name:         name,
		PrepDuration: prep,
		CookDuration: cook,
	}

	error := s.DB.Create(&entry).Error

	return entry, error
}

func (s *RecipeRepository) Get(id int) (models.Recipe, error) {
	var recipe models.Recipe

	error := s.DB.Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.QuantityUnit").
		Preload("Steps").
		Preload("Categories").
		First(&recipe, id).Error

	return recipe, error
}

func (s *RecipeRepository) GetAll() (r []models.Recipe, err error) {
	var recipes []models.Recipe

	error := s.DB.Find(&recipes).Error

	return recipes, error
}

func (s *RecipeRepository) GetRecent(limit int) (r []models.Recipe, err error) {
	var recipes []models.Recipe

	error := s.DB.Order("id DESC").Limit(limit).Find(&recipes).Error

	return recipes, error
}
