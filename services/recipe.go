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
	CategoryRepository         *db.CategoryRepository
}

type RecipeData struct {
	Name          string
	Steps         []string
	Ingredients   []string
	Quantities    []string
	QuantityUnits []string
	Categories    []string
	PrepDuration  uint
	CookDuration  uint
}

func NewRecipeUOW(database *gorm.DB) RecipeUnitOfWork {
	return RecipeUnitOfWork{
		db:                         database,
		RecipeRepository:           db.NewRecipeRepository(database, "recipes"),
		IngredientRepository:       db.NewIngredientRepository(database, "ingredients"),
		QuantityUnitRepository:     db.NewQuantityUnitRepository(database, "quantity_units"),
		RecipeIngredientRepository: db.NewRecipeIngredientRepository(database, "recipe_ingredients"),
		StepRepository:             db.NewStepRepository(database, "steps"),
		CategoryRepository:         db.NewCategoryRepository(database, "categories"),
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

func (s *RecipeService) Create(data RecipeData) (recipe db.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.CreateRecipe(data.Name, data.PrepDuration, data.CookDuration)

	if err != nil {
		s.Log.Error("Error creating recipe", err)
		return result, err
	}

	for i := range data.Ingredients {
		ingredient, _ := s.UnitOfWork.IngredientRepository.GetOrCreate(data.Ingredients[i])
		unit, _ := s.UnitOfWork.QuantityUnitRepository.GetOrCreate(data.QuantityUnits[i])

		s.UnitOfWork.RecipeIngredientRepository.Create(result.ID, ingredient.ID, unit.ID, data.Quantities[i])
	}

	for i := range data.Steps {
		s.UnitOfWork.StepRepository.Create(data.Steps[i], result.ID)
	}

	if data.Categories[0] != "" {
		for i := range data.Categories {
			s.UnitOfWork.CategoryRepository.Create(data.Categories[i], result.ID)
		}
	}

	return result, nil
}

func (s *RecipeService) Update(id int, data RecipeData) (db.Recipe, error) {
	recipe, err := s.UnitOfWork.RecipeRepository.GetRecipe(id)

	if err != nil {
		s.Log.Error("Recipe does not exist", err)
		return db.Recipe{}, err
	}

	for i := range recipe.RecipeIngredients {
		err := s.UnitOfWork.RecipeIngredientRepository.Delete(recipe.RecipeIngredients[i].ID)

		if err != nil {
			s.Log.Error("Error deleting recipe ingredient", err)
			return db.Recipe{}, err
		}
	}

	for i := range data.Ingredients {
		ingredient, _ := s.UnitOfWork.IngredientRepository.GetOrCreate(data.Ingredients[i])
		unit, _ := s.UnitOfWork.QuantityUnitRepository.GetOrCreate(data.QuantityUnits[i])

		s.UnitOfWork.RecipeIngredientRepository.Create(recipe.ID, ingredient.ID, unit.ID, data.Quantities[i])
	}

	for i := range recipe.Steps {
		s.UnitOfWork.StepRepository.Delete(recipe.Steps[i].ID)
	}

	for i := range data.Steps {
		s.UnitOfWork.StepRepository.Create(data.Steps[i], recipe.ID)
	}

	if recipe.Categories != nil {
		for i := range recipe.Categories {
			s.UnitOfWork.CategoryRepository.Delete(recipe.Categories[i].ID)
		}
	}

	if data.Categories[0] != "" {
		for i := range data.Categories {
			s.UnitOfWork.CategoryRepository.Create(data.Categories[i], recipe.ID)
		}
	}

	if data.PrepDuration != 0 {
		recipe.PrepDuration = data.PrepDuration
	}

	if data.CookDuration != 0 {
		recipe.CookDuration = data.CookDuration
	}

	recipe.Name = data.Name

	s.UnitOfWork.db.Save(&recipe)
	recipe, _ = s.UnitOfWork.RecipeRepository.GetRecipe(id)

	return recipe, nil
}
