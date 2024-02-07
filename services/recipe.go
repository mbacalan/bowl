package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/repositories"
	"gorm.io/gorm"
)

type RecipeService struct {
	Log        *slog.Logger
	UnitOfWork RecipeUnitOfWork
}

type RecipeUnitOfWork struct {
	db                         *gorm.DB
	RecipeRepository           *db.RecipeRepository
	IngredientRepository       *db.IngredientRepository
	QuantityUnitRepository     *db.QuantityUnitRepository
	RecipeIngredientRepository *db.RecipeIngredientRepository
	StepRepository             *db.StepRepository
}

func NewRecipeUOW(database *gorm.DB) RecipeUnitOfWork {
	return RecipeUnitOfWork{
		db:                         database,
		RecipeRepository:           db.NewRecipeRepository(database, "recipes"),
		IngredientRepository:       db.NewIngredientRepository(database, "ingredients"),
		QuantityUnitRepository:     db.NewQuantityUnitRepository(database, "quantity_units"),
		RecipeIngredientRepository: db.NewRecipeIngredientRepository(database, "recipe_ingredients"),
		StepRepository:             db.NewStepRepository(database, "steps"),
	}
}

func NewRecipeService(log *slog.Logger, uow RecipeUnitOfWork) RecipeService {
	return RecipeService{
		Log:        log,
		UnitOfWork: uow,
	}
}

func (s *RecipeService) Get(id int) (recipe db.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.GetRecipe(id)

	if err != nil {
		s.Log.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetAll() (recipes []db.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.GetAllRecipes()

	if err != nil {
		s.Log.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetRecent(limit int) (recipes []db.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.GetRecentRecipes(limit)

	if err != nil {
		s.Log.Error("Error getting recent recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) Create(name string, steps []string, ingredients []string, quantities []string, quantityUnits []string) (recipe db.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.CreateRecipe(name)

	for i := range ingredients {
		ingredient, _ := s.UnitOfWork.IngredientRepository.GetOrCreate(ingredients[i])
		unit, _ := s.UnitOfWork.QuantityUnitRepository.GetOrCreate(quantityUnits[i])

		s.UnitOfWork.RecipeIngredientRepository.Create(result.ID, ingredient.ID, unit.ID, quantities[i])
		s.UnitOfWork.StepRepository.Create(steps[i], result.ID)
	}

	if err != nil {
		s.Log.Error("Error creating recipe", err)
		return result, err
	}

	return result, nil
}
