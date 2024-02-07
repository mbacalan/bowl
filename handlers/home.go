package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type HomeHandler struct {
	Log         *slog.Logger
	HomeService services.RecipeService
}

func NewHomeHandler(log *slog.Logger, service services.RecipeService) *HomeHandler {
	return &HomeHandler{
		Log:         log,
		HomeService: service,
	}
}

func (h *HomeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewHome)

	return r
}

func (h *HomeHandler) ViewHome(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.HomeService.GetRecent(10)

	if err != nil {
		h.Log.Error("Error viewing home", err)
	}

	pages.Home(recipes).Render(r.Context(), w)
}
