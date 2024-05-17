package models

import "gorm.io/gorm"

type Repositories struct {
	UserRepository         UserRepository
	RecipeRepository       RecipeUnitOfWork
	IngredientRepository   IngredientRepository
	QuantityUnitRepository QuantityUnitRepository
	CategoryRepository     CategoryRepository
}

type RecipeUnitOfWork struct {
	DB                         *gorm.DB
	RecipeRepository           RecipeRepository
	IngredientRepository       IngredientRepository
	QuantityUnitRepository     QuantityUnitRepository
	RecipeIngredientRepository RecipeIngredientRepository
	StepRepository             StepRepository
	CategoryRepository         CategoryRepository
}

type Repository struct {
	DB        *gorm.DB
	TableName string
}

type CategoryRepository interface {
	Get(id int) (Category, error)
	GetAll() ([]Category, error)
	Delete(id uint) error
}

type UserRepository interface {
	Create(name string, password []byte) (User, error)
	Get(name string) (User, error)
}

type RecipeRepository interface {
	Create(name string, prep uint, cook uint) (Recipe, error)
	Get(id int) (Recipe, error)
	GetAll() ([]Recipe, error)
	GetRecent(limit int) ([]Recipe, error)
	// Delete
}

type RecipeIngredientRepository interface {
	Create(recipeID uint, ingredientID uint, unitID uint, quantity string) (RecipeIngredient, error)
	GetAll() ([]RecipeIngredient, error)
	Delete(id uint) error
}

type IngredientRepository interface {
	GetOrCreate(ingredient string) (Ingredient, error)
	Get(id int) (Ingredient, error)
	GetAll() ([]Ingredient, error)
	Search(name string) ([]Ingredient, error)
}

type QuantityUnitRepository interface {
	GetOrCreate(unit string) (QuantityUnit, error)
	GetAll() ([]QuantityUnit, error)
}

type StepRepository interface {
	Create(step string, recipe uint) (Step, error)
	GetAll() ([]Step, error)
	Delete(id uint) error
}
