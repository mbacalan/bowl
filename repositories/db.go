package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Ingredient{}, &QuantityUnit{}, &RecipeIngredient{}, &Step{}, &Category{}, &Recipe{})

	CreateIfNotExists(db, QuantityUnit{Name: "g"})
	CreateIfNotExists(db, QuantityUnit{Name: "kg"})
	CreateIfNotExists(db, QuantityUnit{Name: "ml"})
	CreateIfNotExists(db, QuantityUnit{Name: "L"})

	return db
}

func CreateIfNotExists(db *gorm.DB, data QuantityUnit) (QuantityUnit, error) {
	var result QuantityUnit

	if err := db.Where(data).FirstOrCreate(&result).Error; err != nil {
		return QuantityUnit{}, err
	}

	return result, nil
}
