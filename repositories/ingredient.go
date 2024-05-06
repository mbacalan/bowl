package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

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

func (s IngredientRepository) GetOrCreate(ingredient string) (models.Ingredient, error) {
	var entry models.Ingredient

	error := s.db.FirstOrCreate(&entry, models.Ingredient{Name: ingredient}).Error

	return entry, error
}

func (s IngredientRepository) Get(id int) (models.Ingredient, error) {
	var ingredient models.Ingredient

	error := s.db.First(&ingredient, id).Error

	return ingredient, error
}

func (s IngredientRepository) GetAll() ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	error := s.db.Find(&ingredients).Error

	return ingredients, error
}

func (s IngredientRepository) Search(name string) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	error := s.db.Where("name LIKE ?", "%"+name+"%").Find(&ingredients).Error

	return ingredients, error
}
