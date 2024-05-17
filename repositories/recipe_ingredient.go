package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type RecipeIngredientRepository struct {
	*models.Repository
}

func NewRecipeIngredientRepository(db *gorm.DB, tableName string) *RecipeIngredientRepository {
	return &RecipeIngredientRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *RecipeIngredientRepository) Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (models.RecipeIngredient, error) {
	entry := models.RecipeIngredient{
		RecipeID:       recipeID,
		IngredientID:   ingredientID,
		QuantityUnitID: unitID,
		Quantity:       quantity,
	}

	error := s.DB.Create(&entry).Error

	return entry, error
}

func (s *RecipeIngredientRepository) GetAll() ([]models.RecipeIngredient, error) {
	var entries []models.RecipeIngredient

	error := s.DB.Find(&entries).Error

	return entries, error
}

func (s *RecipeIngredientRepository) Delete(id uint) error {
	return s.DB.Delete(&models.RecipeIngredient{}, id).Error
}
