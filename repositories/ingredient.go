package repositories

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

func (s IngredientRepository) GetOrCreate(ingredient string) (Ingredient, error) {
	var entry Ingredient

	error := s.db.FirstOrCreate(&entry, Ingredient{Name: ingredient}).Error

	return entry, error
}

func (s IngredientRepository) Get(id int) (Ingredient, error) {
	var ingredient Ingredient

	error := s.db.First(&ingredient, id).Error

	return ingredient, error
}

func (s IngredientRepository) GetAll() ([]Ingredient, error) {
	var ingredients []Ingredient

	error := s.db.Find(&ingredients).Error

	return ingredients, error
}

func (s IngredientRepository) Search(name string) ([]Ingredient, error) {
	var ingredients []Ingredient

	error := s.db.Where("name LIKE ?", "%"+name+"%").Find(&ingredients).Error

	return ingredients, error
}
