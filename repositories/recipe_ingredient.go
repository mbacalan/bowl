package repositories

import (
	"gorm.io/gorm"
)

type RecipeIngredient struct {
	gorm.Model
	RecipeID       uint
	IngredientID   uint
	Ingredient     Ingredient `gorm:"foreignKey:IngredientID"`
	QuantityUnitID uint
	QuantityUnit   QuantityUnit `gorm:"foreignKey:QuantityUnitID"`
	Quantity       string
}

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

func (s RecipeIngredientRepository) Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (RecipeIngredient, error) {
	entry := RecipeIngredient{
		RecipeID:       recipeID,
		IngredientID:   ingredientID,
		QuantityUnitID: unitID,
		Quantity:       quantity,
	}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s RecipeIngredientRepository) GetAll() ([]RecipeIngredient, error) {
	var entries []RecipeIngredient

	error := s.db.Find(&entries).Error

	return entries, error
}

func (s RecipeIngredientRepository) Delete(id uint) error {
	error := s.db.Delete(&RecipeIngredient{}, id).Error

	return error
}
