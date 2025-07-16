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
	"gorm.io/gorm"
)

func newQuantityUnitList(size int) []models.QuantityUnit {
	if size == 0 {
		return make([]models.QuantityUnit, 0)
	}

	quantityUnits := make([]models.QuantityUnit, size)
	for i := 0; i != size; i++ {
		quantityUnits[i] = models.QuantityUnit{
			Model: gorm.Model{ID: uint(i + 1)},
			Name:  fmt.Sprintf("Quantity Unit %d", i+1),
		}
	}

	return quantityUnits
}

func Test_quantityUnitTemplate(t *testing.T) {
	var testCases = []struct {
		name         string
		model        []models.QuantityUnit
		selectedUnit string
		selectors    []string
		matches      []string
	}{
		{
			name:         "default selection fallback is placeholder",
			model:        newQuantityUnitList(2),
			selectedUnit: "",
			selectors:    []string{`[data-testid="quantity-unit-option-default"][selected]`},
			matches:      []string{"Select"},
		},
		{
			name:         "lists units with no default selection",
			model:        newQuantityUnitList(3),
			selectedUnit: "",
			selectors:    []string{`[data-testid="quantity-unit-option-1"]`, `[data-testid="quantity-unit-option-2"]`, `[data-testid="quantity-unit-option-3"]`},
			matches:      []string{"Quantity Unit 1", "Quantity Unit 2", "Quantity Unit 3"},
		},
		{
			name:         "lists units and sets a default selection",
			model:        newQuantityUnitList(3),
			selectedUnit: "Quantity Unit 2",
			selectors:    []string{`[data-testid="quantity-unit-option-1"]`, `[data-testid="quantity-unit-option-2"][selected]`, `[data-testid="quantity-unit-option-3"]`},
			matches:      []string{"Quantity Unit 1", "Quantity Unit 2", "Quantity Unit 3"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r, w := io.Pipe()
			go func() {
				_ = pages.QuantityUnits(test.model, test.selectedUnit).Render(context.Background(), w)
				_ = w.Close()
			}()
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				t.Fatalf("failed to read template: %v", err)
			}

			for i := 0; i != len(test.selectors); i++ {
				selection := doc.Find(test.selectors[i])
				require.Equal(t, 1, len(selection.Nodes), "unexpected # of matches")
				assert.Equal(t, test.matches[i], selection.Text())
			}
		})
	}
}
