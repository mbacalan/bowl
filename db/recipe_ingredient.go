package db

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

type RecipeIngredientStore struct {
	db        *gorm.DB
	tableName string
}

func NewRecipeIngredientStore(db *gorm.DB, tableName string) *RecipeIngredientStore {
	store := &RecipeIngredientStore{
		tableName: tableName,
		db:        db,
	}

	return store
}

func (s RecipeIngredientStore) Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (i RecipeIngredient, err error) {
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

func (s RecipeIngredientStore) GetAll() (recipeIngredients []RecipeIngredient, err error) {
	var entries []RecipeIngredient
	result := s.db.Find(&entries)

	if result.Error != nil {
		return []RecipeIngredient{}, result.Error
	}

	return entries, nil
}
