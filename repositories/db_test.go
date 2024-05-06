package repositories_test

import (
	"testing"

	"github.com/mbacalan/bowl/models"
	"github.com/mbacalan/bowl/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := repositories.NewConnection(sqlite.Open(":memory:"))
	if err != nil {
		t.Fatalf("error connecting database: %v", err)
	}

	return db
}

func TestCreateRepositories(t *testing.T) {
	db := setupTestDB(t)

	repositories.CreateRepositories(db)
}

func TestNewConnection(t *testing.T) {
	db, err := repositories.NewConnection(sqlite.Open(":memory:"))
	if db == nil {
		t.Errorf("error connecting database")
	}

	db, err = repositories.NewConnection(sqlite.Open("sqlserver://test:fail@localhost:9930"))
	if err == nil {
		t.Errorf("expected error connecting database")
	}
}

func TestCreateIfNotExists(t *testing.T) {
	db := setupTestDB(t)
	expected := models.QuantityUnit{Name: "Test unit"}
	actual, err := repositories.CreateIfNotExists(db, expected)
	if err != nil {
		t.Errorf("error creating QuantityUnit: %v", err)
	}

	if expected.Name != actual.Name {
		t.Errorf("expected name %s, got %s", expected.Name, actual.Name)
	}

	_, err = repositories.CreateIfNotExists(db, expected)
	if err != nil {
		t.Errorf("error creating existing QuantityUnit: %v", err)
	}
}

func TestMigrate(t *testing.T) {
	db := setupTestDB(t)

	repositories.Migrate(db)
}

func TestSeedQuantityUnits(t *testing.T) {
	db := setupTestDB(t)

	repositories.SeedQuantityUnits(db)
}
