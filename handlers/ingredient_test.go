package handlers_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/mbacalan/bowl/handlers"
	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services"
)

func TestIngredientHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		assert         func(doc *goquery.Document)
	}{
		// Add error test
		{
			name:           "list ingredients",
			expectedStatus: http.StatusOK,
			assert: func(doc *goquery.Document) {
				if doc.Find(`[data-testid="ingredient-list"]`).Length() == 0 {
					t.Error("expected ingredient to be rendered, but it wasn't")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			database := repositories.NewConnection()
			logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
			repository := repositories.NewIngredientRepository(database, "ingredients")
			service := services.NewIngredientService(logger, repository)

			ih := handlers.NewIngredientHandler(logger, service)

			repository.GetOrCreate("test ingredient")

			ih.Routes().ServeHTTP(w, r)

			doc, err := goquery.NewDocumentFromReader(w.Result().Body)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}

			if test.expectedStatus != w.Code {
				t.Errorf("expected status %d, got %d", test.expectedStatus, w.Code)
			}

			test.assert(doc)
		})
	}
}
