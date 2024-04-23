package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type IngredientHandler struct {
	Log               *slog.Logger
	IngredientService *services.IngredientService
}

func NewIngredientHandler(log *slog.Logger, service *services.IngredientService) *IngredientHandler {
	return &IngredientHandler{
		Log:               log,
		IngredientService: service,
	}
}

func (h *IngredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewIngredientList)
	r.Get("/create", h.CreateIngredient)
	r.Post("/create", h.CreateIngredient)
	// r.Post("/search", h.SearchIngredient)

	return r
}

func (h *IngredientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.CreateIngredient(w, r)
		return
	}

	h.ViewIngredientList(w, r)
}

func (h *IngredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateIngredient().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	_, err := h.IngredientService.Create(r.Form.Get("name"))
	if err != nil {
		h.Log.Error("", err)
		return
	}

	ingredients, err := h.IngredientService.GetAll()
	if err != nil {
		h.Log.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}

func (h *IngredientHandler) ViewIngredientList(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.IngredientService.GetAll()
	if err != nil {
		h.Log.Error("Error listing ingredients", err)
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
