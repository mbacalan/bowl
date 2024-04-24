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
	r.Post("/create", h.Create)
	// r.Post("/search", h.SearchIngredient)

	return r
}

func (h *IngredientHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateIngredient().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	_, err := h.Service.Create(r.Form.Get("name"))
	if err != nil {
		h.Logger.Error("", err)
		return
	}

	ingredients, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}

func (h *IngredientHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}

// func (h *IngredientHandler) SearchIngredient(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	ingredients, err := h.IngredientService.Search(r.Form.Get("search"))
// 	if err != nil {
// 		h.Log.Error("Error searching ingredients", err)
// 	}

// 	pages.IngredientSearchList(ingredients).Render(r.Context(), w)
// }
