package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type homeHandler struct {
	*models.HomeHandler
}

func NewHomeHandler(logger *slog.Logger, service models.HomeService) *homeHandler {
	return &homeHandler{
		HomeHandler: &models.HomeHandler{
			Logger:  logger,
			Service: service,
		},
	}
}

func (h *homeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.View)

	return r
}

func (h *homeHandler) View(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.Service.GetRecent(10)

	if err != nil {
		h.Logger.Error("Error viewing home", err)
	}

	pages.Home(recipes).Render(r.Context(), w)
}
