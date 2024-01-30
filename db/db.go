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

	return db
}
