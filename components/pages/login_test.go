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

func Test_loginTemplate(t *testing.T) {
	var testCases = []struct {
		name     string
		selector string
		matches  []string
	}{
		{
			name:     "form has username input",
			selector: `form#login-form input[name="username"]`,
			matches:  []string{""},
		},
		{
			name:     "form has password input",
			selector: `form#login-form input[name="password"][type="password"]`,
			matches:  []string{""},
		},
		{
			name:     "form has submit button",
			selector: `form#login-form button[type="submit"]`,
			matches:  []string{"Login"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.Login().Render(context.Background(), w)
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
