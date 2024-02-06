package db

import (
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string
}

type IngredientStore struct {
	db        *gorm.DB
	tableName string
}

func NewIngredientStore(db *gorm.DB, tableName string) *IngredientStore {
	store := &IngredientStore{
		tableName: tableName,
		db:        db,
	}

	return store
}

func (s IngredientStore) GetOrCreate(ingredient string) (i Ingredient, err error) {
	var entry Ingredient
	result := s.db.FirstOrCreate(&entry, Ingredient{Name: ingredient})

	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return entry, nil
}

func (s IngredientStore) GetIngredient(id int) (i Ingredient, err error) {
	var ingredient Ingredient
	result := s.db.Find(&ingredient, id)

	if result.Error != nil {
		return Ingredient{}, result.Error
	}

	return ingredient, nil
}

func (s IngredientStore) GetAllIngredients() (i []Ingredient, err error) {
	var ingredients []Ingredient
	result := s.db.Find(&ingredients)

	if result.Error != nil {
		return []Ingredient{}, result.Error
	}

	return ingredients, nil
}

func (s IngredientStore) SearchIngredient(name string) (i []Ingredient, err error) {
	var ingredients []Ingredient
	result := s.db.Where("name LIKE ?", "%"+name+"%").Find(&ingredients)

	if result.Error != nil {
		return []Ingredient{}, result.Error
	}

	return ingredients, nil
}
