package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("error connecting database: %v", err)
	}

	err = db.AutoMigrate(&repositories.Ingredient{}, &repositories.QuantityUnit{}, &repositories.RecipeIngredient{}, &repositories.Step{}, &repositories.Category{}, &repositories.Recipe{})
	if err != nil {
		t.Fatalf("error migrating models: %v", err)
	}

	return db
}
