package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type QuantityUnitHandler struct {
	Log     *slog.Logger
	Service *services.QuantityUnitService
}

func NewQuantityUnitHandler(log *slog.Logger, service *services.QuantityUnitService) *QuantityUnitHandler {
	return &QuantityUnitHandler{
		Log:     log,
		Service: service,
	}
}

func (h *QuantityUnitHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.GetAll)

	return r
}

func (h *QuantityUnitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	selected := r.URL.Query().Get("selected")
	unit, err := h.Service.GetAll()

	if err != nil {
		h.Log.Error("", err)
		return
	}

	if selected != "" {
		pages.QuantityUnits(unit, selected).Render(r.Context(), w)
		return
	}

	pages.QuantityUnits(unit, "").Render(r.Context(), w)
}
