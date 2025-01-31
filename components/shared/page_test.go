package shared_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"
	"github.com/mbacalan/bowl/components/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_pageTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		selector string
		matches  []string
	}{
		{
			name:     "includes head tag",
			selector: `head`,
			matches:  []string{""},
		},
		{
			name:     "includes body tag",
			selector: `body`,
			matches:  []string{""},
		},
		{
			name:     "renders children",
			selector: `[data-testid="page-children"]`,
			matches:  []string{"Page Children"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				contents := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					_, err := io.WriteString(w, "<div data-testid=\"page-children\">Page Children</div>")
					return err
				})
				ctx := templ.WithChildren(context.Background(), contents)
				_ = shared.Page().Render(ctx, w)
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
