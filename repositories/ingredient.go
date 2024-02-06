package db

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string
}

type IngredientRepository struct {
	db        *gorm.DB
	tableName string
}

func NewIngredientRepository(db *gorm.DB, tableName string) *IngredientRepository {
	repository := &IngredientRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s IngredientRepository) GetOrCreate(ingredient string) (i Ingredient, err error) {
	var entry Ingredient
	result := s.db.FirstOrCreate(&entry, Ingredient{Name: ingredient})

	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return entry, nil
}

func (s IngredientRepository) GetIngredient(id int) (i Ingredient, err error) {
	var ingredient Ingredient
	result := s.db.Find(&ingredient, id)

	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return ingredient, nil
}

func (s IngredientRepository) GetAllIngredients() (i []Ingredient, err error) {
	var ingredients []Ingredient
	result := s.db.Find(&ingredients)

	if result.Error != nil {
		return []Ingredient{}, result.Error
	}

	return ingredients, nil
}

func (s IngredientRepository) SearchIngredient(name string) (i []Ingredient, err error) {
	var ingredients []Ingredient
	result := s.db.Where("name LIKE ?", "%"+name+"%").Find(&ingredients)

	if result.Error != nil {
		return []Ingredient{}, result.Error
	}

	return ingredients, nil
}
