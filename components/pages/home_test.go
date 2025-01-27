package pages_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewRecipeList() []models.Recipe {
	return make([]models.Recipe, 0)
}

func AddRecipe(recipes []models.Recipe, name string) []models.Recipe {
	recipe := models.Recipe{
		Name:   name,
		Steps:  make([]models.Step, 0),
		UserID: 1,
	}
	return append(recipes, recipe)
}

func AddRecipes(recipes []models.Recipe, names ...string) []models.Recipe {
	for _, name := range names {
		recipes = AddRecipe(recipes, name)
	}
	return recipes
}

func Test_homeTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		model    []models.Recipe
		path     string
		selector string
		matches  []string
	}{
		{
			name:     "links to recipes",
			model:    AddRecipe(NewRecipeList(), "Test Recipe"),
			selector: "a.recipes-link",
			matches:  []string{"📃 Recipes"},
		},
		{
			name:     "links to categories",
			model:    AddRecipe(NewRecipeList(), "Test Recipe"),
			selector: "a.categories-link",
			matches:  []string{"📚 Categories"},
		},
		{
			name:     "lists recent recipes",
			model:    AddRecipes(NewRecipeList(), "Test Recipe", "Super Soup"),
			selector: "ul.recipe-list a",
			matches:  []string{"Test Recipe", "Super Soup"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.Home(test.model).Render(context.Background(), w)
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
