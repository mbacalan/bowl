package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type Repositories struct {
	RecipeRepository       *RecipeUnitOfWork
	IngredientRepository   *IngredientRepository
	QuantityUnitRepository *QuantityUnitRepository
	CategoryRepository     *CategoryRepository
}

func CreateRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		RecipeRepository:       NewRecipeUOW(db),
		IngredientRepository:   NewIngredientRepository(db, "ingredients"),
		QuantityUnitRepository: NewQuantityUnitRepository(db, "quantity_units"),
		CategoryRepository:     NewCategoryRepository(db, "categories"),
	}
}

func NewConnection(dialector gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Printf("error connecting database: %v", err)
		return nil, err
	}

	Migrate(db)

	return db, nil
}

func CreateIfNotExists(db *gorm.DB, data models.QuantityUnit) (models.QuantityUnit, error) {
	var result models.QuantityUnit

	error := db.Where(data).FirstOrCreate(&result).Error

	return result, error
}

func Migrate(db *gorm.DB) {
	dbModels := []interface{}{
		&models.Ingredient{},
		&models.QuantityUnit{},
		&models.RecipeIngredient{},
		&models.Step{},
		&models.Category{},
		&models.Recipe{},
	}

	for _, model := range dbModels {
		db.AutoMigrate(model)
	}
}

func SeedQuantityUnits(db *gorm.DB) {
	quantityUnitNames := []string{"g", "kg", "ml", "L"}

	for _, name := range quantityUnitNames {
		CreateIfNotExists(db, models.QuantityUnit{Name: name})
	}
}
