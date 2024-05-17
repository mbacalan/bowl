package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type ingredientHandler struct {
	*models.IngredientHandler
}

func NewIngredientHandler(logger *slog.Logger, service models.IngredientService) *ingredientHandler {
	return &ingredientHandler{
		IngredientHandler: &models.IngredientHandler{
			Logger:  logger,
			Service: service,
		},
	}
}

func (h *ingredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)

	return r
}

func (h *ingredientHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}
