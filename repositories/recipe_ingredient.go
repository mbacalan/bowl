package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

type RecipeIngredientRepository struct {
	db        *gorm.DB
	tableName string
}

func NewRecipeIngredientRepository(db *gorm.DB, tableName string) *RecipeIngredientRepository {
	repository := &RecipeIngredientRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s RecipeIngredientRepository) Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (models.RecipeIngredient, error) {
	entry := models.RecipeIngredient{
		RecipeID:       recipeID,
		IngredientID:   ingredientID,
		QuantityUnitID: unitID,
		Quantity:       quantity,
	}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s RecipeIngredientRepository) GetAll() ([]models.RecipeIngredient, error) {
	var entries []models.RecipeIngredient

	error := s.db.Find(&entries).Error

	return entries, error
}

func (s RecipeIngredientRepository) Delete(id uint) error {
	error := s.db.Delete(&models.RecipeIngredient{}, id).Error

	return error
}
