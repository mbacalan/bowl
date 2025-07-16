package pages_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_createIngredientTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		selector string
		matches  []string
	}{
		{
			name:     "form has name input",
			selector: `form#ingredient-form input[name="name"]`,
			matches:  []string{""},
		},
		{
			name:     "form has submit button",
			selector: `form#ingredient-form button[type="submit"]`,
			matches:  []string{"Submit"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.CreateIngredient().Render(context.Background(), w)
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
