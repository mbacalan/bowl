package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type IngredientHandler struct {
	Logger  *slog.Logger
	Service *services.IngredientService
}

func NewIngredientHandler(logger *slog.Logger, service *services.IngredientService) *IngredientHandler {
	return &IngredientHandler{
		Logger:  logger,
		Service: service,
	}
}

func (h *IngredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)

	return r
}

func (h *IngredientHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}
