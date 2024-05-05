package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/repositories"
)

type QuantityUnitHandler struct {
	Logger  *slog.Logger
	Service QuantityUnitService
}

type QuantityUnitService interface {
	GetAll() ([]repositories.QuantityUnit, error)
}

func NewQuantityUnitHandler(logger *slog.Logger, service QuantityUnitService) *QuantityUnitHandler {
	return &QuantityUnitHandler{
		Logger:  logger,
		Service: service,
	}
}

func (h *QuantityUnitHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)

	return r
}

func (h *QuantityUnitHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	selected := r.URL.Query().Get("selected")
	unit, err := h.Service.GetAll()

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	if selected != "" {
		pages.QuantityUnits(unit, selected).Render(r.Context(), w)
		return
	}

	pages.QuantityUnits(unit, "").Render(r.Context(), w)
}
