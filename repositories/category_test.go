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

	err = db.AutoMigrate(&repositories.Category{}, &repositories.Recipe{})
	if err != nil {
		t.Fatalf("error migrating models: %v", err)
	}

	return db
}

func TestCategoryRepository(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := repositories.Category{Name: "Test Category"}

		db.Create(&expected)
		actual, err := repo.Get(int(expected.ID))

		if err != nil {
			t.Errorf("error getting category: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected category name %s, got %s", expected.Name, actual.Name)
		}

		if actual.Recipes == nil {
			t.Errorf("expected non-nil recipes for category %s", actual.Name)
		}
	})

	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := []repositories.Category{
			{Name: "Test Category"},
			{Name: "Test Category 2"},
			{Name: "Test Category 3"},
		}

		db.Create(&expected)
		actual, err := repo.GetAll()

		if err != nil {
			t.Errorf("error getting all categories: %v", err)
		}

		var categories []repositories.Category
		var count int64
		db.Find(&categories).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d categories, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected category name %s, got %s", expected[i].Name, actual[i].Name)
			}

			if actual[i].Recipes == nil {
				t.Errorf("expected non-nil recipes for category %s", actual[i].Name)
			}
		}
	})

	t.Run("delete", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := []repositories.Category{
			{Name: "Test Category"},
			{Name: "Test Category 2"},
			{Name: "Test Category 3"},
		}

		db.Create(&expected)

		var categories []repositories.Category
		db.Find(&categories)

		err := repo.Delete(categories[0].ID)
		if err != nil {
			t.Errorf("error deleting category: %v", err)
		}

		var count int64
		db.Find(&categories).Count(&count)

		expectedCount := int64(len(expected) - 1)
		if count != expectedCount {
			t.Errorf("expected %d categories, got %d", expectedCount, count)
		}
	})
}
