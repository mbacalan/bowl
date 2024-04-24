package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type HomeHandler struct {
	Logger  *slog.Logger
	Service *services.RecipeService
}

func NewHomeHandler(logger *slog.Logger, service *services.RecipeService) *HomeHandler {
	return &HomeHandler{
		Logger:  logger,
		Service: service,
	}
}

func (h *HomeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.View)

	return r
}

func (h *HomeHandler) View(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.Service.GetRecent(10)

	if err != nil {
		h.Logger.Error("Error viewing home", err)
	}

	pages.Home(recipes).Render(r.Context(), w)
}
