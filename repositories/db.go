package repositories

import (
	"fmt"

	"gorm.io/gorm"
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

	db.AutoMigrate(&Ingredient{}, &QuantityUnit{}, &RecipeIngredient{}, &Step{}, &Category{}, &Recipe{})

	return db, nil
}

func CreateIfNotExists(db *gorm.DB, data QuantityUnit) (QuantityUnit, error) {
	var result QuantityUnit

	error := db.Where(data).FirstOrCreate(&result).Error

	return result, error
}

func SeedQuantityUnits(db *gorm.DB) {
	quantityUnitNames := []string{"g", "kg", "ml", "L"}

	for _, name := range quantityUnitNames {
		CreateIfNotExists(db, QuantityUnit{Name: name})
	}
}
