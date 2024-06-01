package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password []byte
}

type Recipe struct {
	gorm.Model
	Name              string
	RecipeIngredients []RecipeIngredient
	Steps             []Step
	PrepDuration      uint
	CookDuration      uint
	Categories        []*Category `gorm:"many2many:recipe_categories;"`
	UserID            uint
	// Rating      uint
	// Difficulty  uint
	// WwPoints    uint `gorm:"column:ww_points"`
	// Language    string
	// Comments    []string
}

type Category struct {
	gorm.Model
	Name    string
	Recipes []*Recipe `gorm:"many2many:recipe_categories;"`
}

type Ingredient struct {
	gorm.Model
	Name string
}

type RecipeIngredient struct {
	gorm.Model
	RecipeID       uint
	IngredientID   uint
	Ingredient     Ingredient `gorm:"foreignKey:IngredientID"`
	QuantityUnitID uint
	QuantityUnit   QuantityUnit `gorm:"foreignKey:QuantityUnitID"`
	Quantity       string
}

type QuantityUnit struct {
	gorm.Model
	Name string
}

type Step struct {
	gorm.Model
	RecipeID uint
	Step     string
}
