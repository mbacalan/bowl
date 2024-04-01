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

func (s RecipeIngredientRepository) Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (i RecipeIngredient, err error) {
	entry := RecipeIngredient{
		RecipeID:       recipeID,
		IngredientID:   ingredientID,
		QuantityUnitID: unitID,
		Quantity:       quantity,
	}

	result := s.db.Create(&entry)

	if result.Error != nil {
		return RecipeIngredient{}, result.Error
	}

	return entry, nil
}

func (s RecipeIngredientRepository) GetAll() (recipeIngredients []RecipeIngredient, err error) {
	var entries []RecipeIngredient
	result := s.db.Find(&entries)

	if result.Error != nil {
		return []RecipeIngredient{}, result.Error
	}

	return entries, nil
}

func (s RecipeIngredientRepository) Delete(id uint) (err error) {
	result := s.db.Delete(&RecipeIngredient{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
