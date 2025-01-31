package partials_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/components/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_headerTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		selector string
		matches  []string
	}{
		{
			name:     "includes header tag",
			selector: `header`,
			matches:  []string{""},
		},
		{
			name:     "links to home page",
			selector: `a[data-testid="home-page-link"]`,
			matches:  []string{""},
		},
		{
			name:     "links to recipe create",
			selector: `a[data-testid="recipe-create-link"]`,
			matches:  []string{""},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = shared.Page().Render(context.Background(), w)
				_ = w.Close()
			}()
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}
			selection := doc.Find(test.selector)
			require.Equal(t, len(test.matches), len(selection.Nodes), "unexpected # of matches")
			selection.Each(func(i int, s *goquery.Selection) {
				if test.matches[i] == "" {
					return
				}

				assert.Equal(t, test.matches[i], s.Text())
			})
		})
	}
}
