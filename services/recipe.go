package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type RecipeService struct {
	Logger     *slog.Logger
	UnitOfWork models.RecipeUnitOfWork
}

func NewRecipeService(log *slog.Logger, uow models.RecipeUnitOfWork) *RecipeService {
	return &RecipeService{
		Logger:     log,
		UnitOfWork: uow,
	}
}

func (s *RecipeService) Get(id int) (recipe models.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.Get(id)

	if err != nil {
		s.Logger.Error("Error getting recipe", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetAll() (recipes []models.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.GetAll()

	if err != nil {
		s.Logger.Error("Error getting all recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) GetRecent(limit int) (recipes []models.Recipe, error error) {
	result, err := s.UnitOfWork.RecipeRepository.GetRecent(limit)

	if err != nil {
		s.Logger.Error("Error getting recent recipes", err)
		return result, err
	}

	return result, nil
}

func (s *RecipeService) Create(data models.RecipeData) (models.Recipe, error) {
	recipe, err := s.UnitOfWork.RecipeRepository.Create(data.Name, data.PrepDuration, data.CookDuration)

	if err != nil {
		s.Logger.Error("Error creating recipe", err)
		return recipe, err
	}

	s.createRecipeIngredients(recipe.ID, data.Ingredients, data.Quantities, data.QuantityUnits)
	s.createSteps(recipe.ID, data.Steps)

	if data.Categories[0] != "" {
		for _, categoryName := range data.Categories {
			var category models.Category
			error := s.UnitOfWork.DB.FirstOrCreate(&category, models.Category{Name: cases.Title(language.English).String(categoryName)}).Error

			if error == nil {
				s.UnitOfWork.DB.Model(&recipe).Association("Categories").Append(&category)
			}
		}
	}

	return recipe, nil
}

func (s *RecipeService) Update(id int, data models.RecipeData) (models.Recipe, error) {
	recipe, err := s.UnitOfWork.RecipeRepository.Get(id)

	if err != nil {
		s.Logger.Error("Recipe does not exist", err)
		return models.Recipe{}, err
	}

	for i := range recipe.RecipeIngredients {
		err := s.UnitOfWork.RecipeIngredientRepository.Delete(recipe.RecipeIngredients[i].ID)

		if err != nil {
			s.Logger.Error("Error deleting recipe ingredient", err)
			return models.Recipe{}, err
		}
	}

	s.createRecipeIngredients(recipe.ID, data.Ingredients, data.Quantities, data.QuantityUnits)

	for i := range recipe.Steps {
		s.UnitOfWork.StepRepository.Delete(recipe.Steps[i].ID)
	}

	s.createSteps(recipe.ID, data.Steps)

	if recipe.Categories != nil {
		for i := range recipe.Categories {
			s.UnitOfWork.CategoryRepository.Delete(recipe.Categories[i].ID)
		}
	}

	if data.Categories[0] != "" {
		for i := range data.Categories {
			var category models.Category
			error := s.UnitOfWork.DB.Find(&category, "name = ?", data.Categories[i]).Error

			if error == nil {
				s.UnitOfWork.DB.Model(&recipe).Association("Categories").Append(&models.Category{Name: cases.Title(language.English).String(data.Categories[i])})
			}
		}
	}

	if data.PrepDuration != 0 {
		recipe.PrepDuration = data.PrepDuration
	}

	if data.CookDuration != 0 {
		recipe.CookDuration = data.CookDuration
	}

	recipe.Name = data.Name

	s.UnitOfWork.DB.Save(&recipe)
	recipe, _ = s.UnitOfWork.RecipeRepository.Get(id)

	return recipe, nil
}

func (s *RecipeService) createRecipeIngredients(recipeID uint, ingredients []string, quantities []string, quantityUnits []string) {
	for i := range ingredients {
		ingredient, _ := s.UnitOfWork.IngredientRepository.GetOrCreate(cases.Title(language.English).String(ingredients[i]))
		unit, _ := s.UnitOfWork.QuantityUnitRepository.GetOrCreate(quantityUnits[i])

		s.UnitOfWork.RecipeIngredientRepository.Create(recipeID, ingredient.ID, unit.ID, quantities[i])
	}
}

func (s *RecipeService) createSteps(recipeID uint, steps []string) {
	for i := range steps {
		s.UnitOfWork.StepRepository.Create(steps[i], recipeID)
	}
}
