package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
)

func TestRecipeRepository(t *testing.T) {
	t.Run("get", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeRepository(db, "recipes")
		expected := models.Recipe{Name: "Test Recipe", UserID: 1}

		db.Create(&expected)
		actual, err := repo.Get(1, int(expected.ID))

		if err != nil {
			t.Errorf("error getting recipe: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected recipe name %s, got %s", expected.Name, actual.Name)
		}
	})

	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeRepository(db, "recipes")
		expected := []models.Recipe{
			{Name: "Test Recipe 1", UserID: 1},
			{Name: "Test Recipe 2", UserID: 1},
			{Name: "Test Recipe 3", UserID: 1},
		}

		db.Create(&expected)
		actual, err := repo.GetAll(1)

		if err != nil {
			t.Errorf("error getting all recipes: %v", err)
		}

		var recipes []models.Recipe
		var count int64
		db.Find(&recipes).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d recipes, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected recipe name %s, got %s", expected[i].Name, actual[i].Name)
			}
		}
	})

	t.Run("create", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeRepository(db, "recipes")
		expected := models.Recipe{
			Name:         "Test Recipe",
			PrepDuration: 1,
			CookDuration: 2,
		}

		db.Create(&expected)
		actual, err := repo.Create(expected.Name, expected.PrepDuration, expected.CookDuration, 1)

		if err != nil {
			t.Errorf("error creating recipe: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected recipe %s, got %s", expected.Name, actual.Name)
		}

		if expected.PrepDuration != actual.PrepDuration {
			t.Errorf("expected recipe prep duration %d, got %d", expected.PrepDuration, actual.PrepDuration)
		}

		if expected.CookDuration != actual.CookDuration {
			t.Errorf("expected recipe cook duration %d, got %d", expected.CookDuration, actual.CookDuration)
		}
	})

	t.Run("get recent", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeRepository(db, "recipes")
		expected := []models.Recipe{
			{Name: "Test Recipe 1", UserID: 1},
			{Name: "Test Recipe 2", UserID: 1},
			{Name: "Test Recipe 3", UserID: 1},
		}

		db.Create(&expected)
		actual, err := repo.GetRecent(1, 2)

		if err != nil {
			t.Errorf("error getting recent recipes: %v", err)
		}

		for i := 0; i < 2; i++ {
			if actual[i].Name != expected[2-i].Name {
				t.Errorf("expected recipe %s, but got %s", expected[2-i].Name, actual[i].Name)
			}
		}
	})
}
