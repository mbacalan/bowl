package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

type RecipeRepository struct {
	db        *gorm.DB
	tableName string
}

type RecipeUnitOfWork struct {
	DB                         *gorm.DB
	RecipeRepository           *RecipeRepository
	IngredientRepository       *IngredientRepository
	QuantityUnitRepository     *QuantityUnitRepository
	RecipeIngredientRepository *RecipeIngredientRepository
	StepRepository             *StepRepository
	CategoryRepository         *CategoryRepository
}

func NewRecipeRepository(db *gorm.DB, tableName string) *RecipeRepository {
	repository := &RecipeRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func NewRecipeUOW(database *gorm.DB) *RecipeUnitOfWork {
	return &RecipeUnitOfWork{
		DB:                         database,
		RecipeRepository:           NewRecipeRepository(database, "recipes"),
		IngredientRepository:       NewIngredientRepository(database, "ingredients"),
		QuantityUnitRepository:     NewQuantityUnitRepository(database, "quantity_units"),
		RecipeIngredientRepository: NewRecipeIngredientRepository(database, "recipe_ingredients"),
		StepRepository:             NewStepRepository(database, "steps"),
		CategoryRepository:         NewCategoryRepository(database, "categories"),
	}
}

func (s RecipeRepository) Create(name string, prep uint, cook uint) (models.Recipe, error) {
	entry := models.Recipe{
		Name:         name,
		PrepDuration: prep,
		CookDuration: cook,
	}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s RecipeRepository) Get(id int) (models.Recipe, error) {
	var recipe models.Recipe

	error := s.db.Preload("RecipeIngredients").
		Preload("RecipeIngredients.Ingredient").
		Preload("RecipeIngredients.QuantityUnit").
		Preload("Steps").
		Preload("Categories").
		First(&recipe, id).Error

	return recipe, error
}

func (s RecipeRepository) GetAll() (r []models.Recipe, err error) {
	var recipes []models.Recipe

	error := s.db.Find(&recipes).Error

	return recipes, error
}

func (s RecipeRepository) GetRecent(limit int) (r []models.Recipe, err error) {
	var recipes []models.Recipe

	error := s.db.Order("id DESC").Limit(limit).Find(&recipes).Error

	return recipes, error
}
