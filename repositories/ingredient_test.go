package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/repositories"
)

func TestIngredientRepository(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewIngredientRepository(db, "ingredients")
		expected := repositories.Ingredient{Name: "Test Ingredient"}

		db.Create(&expected)

		if _, err := repo.Get(1337); err == nil {
			t.Errorf("expected error getting non-existent ingredient, got nil")
		}

		actual, err := repo.Get(int(expected.ID))

		if err != nil {
			t.Errorf("error getting ingredient: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected ingredient name %s, got %s", expected.Name, actual.Name)
		}
	})

	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewIngredientRepository(db, "ingredients")
		expected := []repositories.Ingredient{
			{Name: "Test Ingredient 1"},
			{Name: "Test Ingredient 2"},
			{Name: "Test Ingredient 3"},
		}

		db.Create(&expected)
		actual, err := repo.GetAll()

		if err != nil {
			t.Errorf("error getting all ingredients: %v", err)
		}

		var ingredients []repositories.Ingredient
		var count int64
		db.Find(&ingredients).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d ingredients, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected ingredients name %s, got %s", expected[i].Name, actual[i].Name)
			}
		}
	})

	t.Run("get or create", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewIngredientRepository(db, "ingredients")
		expected := repositories.Ingredient{Name: "Test Ingredient"}

		// Get
		db.Create(&expected)
		actual, err := repo.GetOrCreate(expected.Name)

		if err != nil {
			t.Errorf("error creating ingredient: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected ingredient name %s, got %s", expected.Name, actual.Name)
		}

		// Create
		expected = repositories.Ingredient{Name: "Test Create Ingredient"}
		actual, err = repo.GetOrCreate(expected.Name)

		if err != nil {
			t.Errorf("error creating ingredient: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected ingredient name %s, got %s", expected.Name, actual.Name)
		}
	})

	t.Run("search", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewIngredientRepository(db, "ingredients")
		expected := []repositories.Ingredient{
			{Name: "Test Ingredient 1"},
			{Name: "Test Ingredient 2"},
			{Name: "Wont match"},
		}

		db.Create(&expected)
		actual, err := repo.Search("Test Ingredient")

		if err != nil {
			t.Errorf("error searching ingredient: %v", err)
		}

		if len(actual) != 2 {
			t.Errorf("expected %d ingredients, got %d", 2, len(actual))
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected ingredient name %s, got %s", expected[i].Name, actual[i].Name)
			}
		}
	})
}
