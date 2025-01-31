package partials_test

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/components/partials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_headTemplate(t *testing.T) {
	var testCases = []struct {
		name      string
		selectors []string
		matches   []string
	}{
		{
			name:      "includes head tag",
			selectors: []string{`head`},
			matches:   []string{""},
		},
		{
			name:      "sets title",
			selectors: []string{`title`},
			matches:   []string{"🥗 Bowl"},
		},
		{
			name:      "includes htmx",
			selectors: []string{`script[src="/assets/htmx.min.js"]`},
			matches:   []string{""},
		},
		{
			name:      "includes styles",
			selectors: []string{`link[rel="stylesheet"][href="/assets/modern-normalize.css"]`, `link[rel="stylesheet"][href="/assets/sakura.css"]`, `link[rel="stylesheet"][href="assets/style.css"]`},
			matches:   []string{""},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = partials.Head().Render(context.Background(), w)
				_ = w.Close()
			}()
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}
			for i := 0; i != len(test.selectors); i++ {
				selection := doc.Find(test.selectors[i])
				require.Equal(t, 1, len(selection.Nodes), "unexpected # of matches")
				if test.matches[i] == "" {
					return
				}

				assert.Equal(t, test.matches[i], selection.Text())
			}
		})
	}
}
