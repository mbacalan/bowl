package handlers

import (
	"log/slog"
	"net/http"
	// "strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type RecipeIngredientHandler struct {
	Log               *slog.Logger
	IngredientService *services.RecipeIngredientService
}

func NewRecipeIngredientHandler(log *slog.Logger, service *services.RecipeIngredientService) *RecipeIngredientHandler {
	return &RecipeIngredientHandler{
		Log:               log,
		IngredientService: service,
	}
}

func (h *RecipeIngredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/", h.ViewIngredientList)
	r.Get("/create", h.CreateRecipeIngredient)
	r.Post("/create", h.CreateRecipeIngredient)

	return r
}

func (h *RecipeIngredientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.CreateRecipeIngredient(w, r)
		return
	}

	// h.ViewIngredientList(w, r)
}

func (h *RecipeIngredientHandler) CreateRecipeIngredient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateIngredient().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	// ingredient, err := h.IngredientService.Create(r.Form.Get("name"))

	// if err != nil {
	// 	h.Log.Error("", err)
	// 	return
	// }

	// w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(ingredient.ID), 10))
	// pages.IngredientDetailPage(ingredient).Render(r.Context(), w)
}

// func (h *RecipeIngredientHandler) ViewIngredientList(w http.ResponseWriter, r *http.Request) {
// 	ingredients, err := h.IngredientService.GetAll()
// 	if err != nil {
// 		h.Log.Error("Error listing ingredients", err)
// 	}

// pages.IngredientListPage(ingredients).Render(r.Context(), w)
// }
