package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&Recipe{})
	db.AutoMigrate(&QuantityUnit{})

	CreateIfNotExists(db, QuantityUnit{Unit: "g"})
	CreateIfNotExists(db, QuantityUnit{Unit: "kg"})
	CreateIfNotExists(db, QuantityUnit{Unit: "ml"})
	CreateIfNotExists(db, QuantityUnit{Unit: "L"})

	return db
}

func CreateIfNotExists(db *gorm.DB, data QuantityUnit) (QuantityUnit, error) {
	var result QuantityUnit

	if err := db.Where(data).FirstOrCreate(&result).Error; err != nil {
		return QuantityUnit{}, err
	}

	return result, nil
}
