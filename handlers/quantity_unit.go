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
	Service services.QuantityUnitService
}

func NewQuantityUnitHandler(log *slog.Logger, service services.QuantityUnitService) *QuantityUnitHandler {
	return &QuantityUnitHandler{
		Log:     log,
		Service: service,
	}
}

func (h *QuantityUnitHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/", h.ViewIngredientList)
	r.Get("/", h.GetAll)
	// r.Post("/create", h.CreateRecipeIngredient)

	return r
}

func (h *QuantityUnitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodPost {
	// 	h.CreateRecipeIngredient(w, r)
	// 	return
	// }

	h.GetAll(w, r)
}

func (h *QuantityUnitHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		r.ParseForm()

		unit, err := h.Service.GetAll()

		if err != nil {
			h.Log.Error("", err)
			return
		}

		pages.QuantityUnits(unit).Render(r.Context(), w)
	}
}
