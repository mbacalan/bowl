package repositories_test

import (
	"strconv"
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
)

func mockStep(id uint) models.Step {
	return models.Step{
		RecipeID: id,
		Step:     "Step " + strconv.Itoa(int(id)),
	}
}

func TestStepRepository(t *testing.T) {
	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewStepRepository(db, "steps")
		expected := []models.Step{
			mockStep(1),
			mockStep(2),
			mockStep(3),
		}

		db.Create(&expected)
		actual, err := repo.GetAll()

		if err != nil {
			t.Errorf("error getting all steps: %v", err)
		}

		var steps []models.Step
		var count int64
		db.Find(&steps).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d steps, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].RecipeID != actual[i].RecipeID {
				t.Errorf("expected recipe id %d, got %d", expected[i].RecipeID, actual[i].RecipeID)
			}

			if expected[i].Step != actual[i].Step {
				t.Errorf("expected step %s, got %s", expected[i].Step, actual[i].Step)
			}
		}
	})

	t.Run("create", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewStepRepository(db, "steps")
		expected := mockStep(1)

		db.Create(&expected)
		actual, err := repo.Create(expected.Step, expected.RecipeID)

		if err != nil {
			t.Errorf("error creating step: %v", err)
		}

		if expected.RecipeID != actual.RecipeID {
			t.Errorf("expected recipe id %d, got %d", expected.RecipeID, actual.RecipeID)
		}

		if expected.Step != actual.Step {
			t.Errorf("expected step %s, got %s", expected.Step, actual.Step)
		}
	})

	t.Run("delete", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewStepRepository(db, "steps")
		expected := []models.Step{
			mockStep(1),
			mockStep(2),
			mockStep(3),
		}

		db.Create(&expected)

		var steps []models.Step
		db.Find(&steps)

		err := repo.Delete(steps[0].ID)
		if err != nil {
			t.Errorf("error deleting steps: %v", err)
		}

		var count int64
		db.Find(&steps).Count(&count)

		expectedCount := int64(len(expected) - 1)
		if count != expectedCount {
			t.Errorf("expected %d steps, got %d", expectedCount, count)
		}
	})
}
