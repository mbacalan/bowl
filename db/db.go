package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string
}

type RecipeIngredient struct {
	gorm.Model
	RecipeID     uint
	IngredientID uint
	Quantity     string
}

type Category struct {
	gorm.Model
	Name     string
	RecipeID uint
}

type Recipe struct {
	gorm.Model
	Name string
	// Ingredients []RecipeIngredient
	// Steps       []string
	// Rating      uint
	// Difficulty  uint
	// Duration    uint
	// Categories  []Category
	// WwPoints    uint `gorm:"column:ww_points"`
	// Language    string
	// Comments    []string
}

type RecipeStore struct {
	db        *gorm.DB
	tableName string
}

func NewRecipeStore(db *gorm.DB, tableName string) *RecipeStore {
	store := &RecipeStore{
		tableName: tableName,
		db:        db,
	}

	return store
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Ingredient{})
	db.AutoMigrate(&Recipe{})

	return db
}

func CreateRecipe(recipe Recipe) (r Recipe, err error) {
	db, err := gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error creating recipe", err)
		return
	}

	entry := Recipe{Name: recipe.Name}
	result := db.Create(&entry)

	if result.Error != nil {
		return Recipe{}, result.Error
	}

	return entry, nil
}

func (s RecipeStore) Get(id int) (r Recipe, err error) {
	var recipe Recipe
	result := s.db.Find(&recipe, id)

	if result.Error != nil {
		return Recipe{}, result.Error
	}

	return recipe, nil
}

func (s RecipeStore) GetAll() (r []Recipe, err error) {
	var recipes []Recipe
	result := s.db.Find(&recipes)

	if result.Error != nil {
		return []Recipe{}, result.Error
	}

	return recipes, nil
}
