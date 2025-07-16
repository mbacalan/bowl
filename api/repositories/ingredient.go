package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type IngredientRepository struct {
	*models.Repository
}

func NewIngredientRepository(db *gorm.DB, tableName string) *IngredientRepository {
	return &IngredientRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *IngredientRepository) GetOrCreate(ingredient string) (models.Ingredient, error) {
	var entry models.Ingredient

	error := s.DB.FirstOrCreate(&entry, models.Ingredient{Name: ingredient}).Error

	return entry, error
}

func (s *IngredientRepository) Get(id int) (models.Ingredient, error) {
	var ingredient models.Ingredient

	error := s.DB.First(&ingredient, id).Error

	return ingredient, error
}

func (s *IngredientRepository) GetAll() ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	error := s.DB.Find(&ingredients).Error

	return ingredients, error
}

func (s *IngredientRepository) Search(name string) ([]models.Ingredient, error) {
	var ingredients []models.Ingredient

	error := s.DB.Where("name LIKE ?", "%"+name+"%").Find(&ingredients).Error

	return ingredients, error
}
