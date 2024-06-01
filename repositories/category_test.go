package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
)

func TestCategoryRepository(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := models.Category{Name: "Test Category"}

		db.Create(&expected)

		if _, err := repo.Get(1, 1337); err == nil {
			t.Errorf("expected error getting non-existent category, got nil")
		}

		actual, err := repo.Get(1, int(expected.ID))

		if err != nil {
			t.Errorf("error getting category: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected category name %s, got %s", expected.Name, actual.Name)
		}
	})

	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := []models.Category{
			{Name: "Test Category"},
			{Name: "Test Category 2"},
			{Name: "Test Category 3"},
		}

		db.Create(&expected)
		actual, err := repo.GetAll(1)

		if err != nil {
			t.Errorf("error getting all categories: %v", err)
		}

		var categories []models.Category
		var count int64
		db.Find(&categories).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d categories, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected category name %s, got %s", expected[i].Name, actual[i].Name)
			}
		}
	})

	t.Run("delete", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewCategoryRepository(db, "categories")
		expected := []models.Category{
			{Name: "Test Category"},
			{Name: "Test Category 2"},
			{Name: "Test Category 3"},
		}

		db.Create(&expected)

		var categories []models.Category
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
