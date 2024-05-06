package repositories_test

import (
	"strconv"
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
)

func mockRecipeIngredient(id uint) models.RecipeIngredient {
	return models.RecipeIngredient{
		RecipeID:       id,
		IngredientID:   id,
		QuantityUnitID: id,
		Quantity:       strconv.Itoa(int(id)),
	}
}

func TestRecipeIngredientRepository(t *testing.T) {
	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeIngredientRepository(db, "recipe_ingredients")
		expected := []models.RecipeIngredient{
			mockRecipeIngredient(1),
			mockRecipeIngredient(2),
			mockRecipeIngredient(3),
		}

		db.Create(&expected)
		actual, err := repo.GetAll()

		if err != nil {
			t.Errorf("error getting all recipe ingredients: %v", err)
		}

		var recipe_ingredients []models.RecipeIngredient
		var count int64
		db.Find(&recipe_ingredients).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d recipe ingredients, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].RecipeID != actual[i].RecipeID {
				t.Errorf("expected recipe id %d, got %d", expected[i].RecipeID, actual[i].RecipeID)
			}
		}
	})

	t.Run("create", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeIngredientRepository(db, "recipe_ingredients")
		expected := mockRecipeIngredient(1)

		db.Create(&expected)
		actual, err := repo.Create(expected.RecipeID, expected.IngredientID, expected.QuantityUnitID, expected.Quantity)

		if err != nil {
			t.Errorf("error creating recipe ingredient: %v", err)
		}

		if expected.RecipeID != actual.RecipeID {
			t.Errorf("expected recipe id %d, got %d", expected.RecipeID, actual.RecipeID)
		}

		if expected.IngredientID != actual.IngredientID {
			t.Errorf("expected ingredient id %d, got %d", expected.IngredientID, actual.IngredientID)
		}

		if expected.QuantityUnitID != actual.QuantityUnitID {
			t.Errorf("expected quantity unit id %d, got %d", expected.QuantityUnitID, actual.QuantityUnitID)
		}

		if expected.Quantity != actual.Quantity {
			t.Errorf("expected quantity %s, got %s", expected.Quantity, actual.Quantity)
		}
	})

	t.Run("delete", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewRecipeIngredientRepository(db, "recipe_ingredients")
		expected := []models.RecipeIngredient{
			mockRecipeIngredient(1),
			mockRecipeIngredient(2),
			mockRecipeIngredient(3),
		}

		db.Create(&expected)

		var recipe_ingredients []models.RecipeIngredient
		db.Find(&recipe_ingredients)

		err := repo.Delete(recipe_ingredients[0].ID)
		if err != nil {
			t.Errorf("error deleting recipe ingredient: %v", err)
		}

		var count int64
		db.Find(&recipe_ingredients).Count(&count)

		expectedCount := int64(len(expected) - 1)
		if count != expectedCount {
			t.Errorf("expected %d recipe ingredients, got %d", expectedCount, count)
		}
	})
}
