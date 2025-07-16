package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
)

func TestQuantityUnitRepository(t *testing.T) {
	t.Run("get or create", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewQuantityUnitRepository(db, "quantity_units")
		expected := models.QuantityUnit{Name: "Test Kilos"}

		// Get
		db.Create(&expected)
		actual, err := repo.GetOrCreate(expected.Name)

		if err != nil {
			t.Errorf("error getting quantity unit: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected quantity unit name %s, got %s", expected.Name, actual.Name)
		}

		// Create
		expected = models.QuantityUnit{Name: "Test Grams"}
		actual, err = repo.GetOrCreate(expected.Name)

		if err != nil {
			t.Errorf("error getting quantity unit: %v", err)
		}

		if expected.Name != actual.Name {
			t.Errorf("expected quantity unit name %s, got %s", expected.Name, actual.Name)
		}
	})

	t.Run("get all", func(t *testing.T) {
		db := setupTestDB(t)
		repo := repositories.NewQuantityUnitRepository(db, "quantity_units")
		expected := []models.QuantityUnit{
			{Name: "Test Unit 1"},
			{Name: "Test Unit 2"},
			{Name: "Test Unit 3"},
		}

		db.Create(&expected)
		actual, err := repo.GetAll()

		if err != nil {
			t.Errorf("error getting all quantity units: %v", err)
		}

		var quantity_units []models.QuantityUnit
		var count int64
		db.Find(&quantity_units).Count(&count)

		if count != int64(len(expected)) {
			t.Errorf("expected %d quantity units, got %d", len(expected), count)
		}

		for i := range actual {
			if expected[i].Name != actual[i].Name {
				t.Errorf("expected quantity unit name %s, got %s", expected[i].Name, actual[i].Name)
			}
		}
	})
}
