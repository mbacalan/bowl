package pages_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newCategoryList(size int) []models.Category {
	if size == 0 {
		return make([]models.Category, 0)
	}

	categories := make([]models.Category, size)
	for i := 0; i != size; i++ {
		categories[i] = models.Category{
			Name:    fmt.Sprintf("Category %d", i),
			Recipes: newCategoryRecipeList(),
		}
	}

	return categories
}

func newCategoryRecipeList() []*models.Recipe {
	categoryRecipes := make([]*models.Recipe, 2)
	for i := 0; i != 2; i++ {
		categoryRecipes[i] = &models.Recipe{
			Name: fmt.Sprintf("Recipe %d", i),
		}
	}
	return categoryRecipes
}

func Test_categoriesTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		model    []models.Category
		selector string
		matches  []string
	}{
		{
			name:     "shows no categories message",
			selector: `[data-testid="no-categories"]`,
			model:    newCategoryList(0),
			matches:  []string{"No categories found :("},
		},
		{
			name:     "renders category list",
			selector: `[data-testid="category-link"]`,
			model:    newCategoryList(2),
			matches:  []string{"Category 0", "Category 1"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.Categories(test.model).Render(context.Background(), w)
				_ = w.Close()
			}()
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}
			selection := doc.Find(test.selector)
			require.Equal(t, len(test.matches), len(selection.Nodes), "unexpected # of matches")
			selection.Each(func(i int, s *goquery.Selection) {
				assert.Equal(t, test.matches[i], s.Text())
			})
		})
	}
}

func newCategory() models.Category {
	return models.Category{
		Name:    "Category 0",
		Recipes: newCategoryRecipeList(),
	}
}

func Test_categoryTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		model    models.Category
		selector string
		matches  []string
	}{
		{
			name:     "renders category recipe list",
			selector: `[data-testid="category-recipe-link"]`,
			model:    newCategory(),
			matches:  []string{"Recipe 0", "Recipe 1"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.Category(test.model).Render(context.Background(), w)
				_ = w.Close()
			}()
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}
			selection := doc.Find(test.selector)
			require.Equal(t, len(test.matches), len(selection.Nodes), "unexpected # of matches")
			selection.Each(func(i int, s *goquery.Selection) {
				assert.Equal(t, test.matches[i], s.Text())
			})
		})
	}
}
