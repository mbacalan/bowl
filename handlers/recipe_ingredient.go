package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type RecipeIngredientHandler struct {
	Logger  *slog.Logger
	Service *services.RecipeIngredientService
}

func NewRecipeIngredientHandler(logger *slog.Logger, service *services.RecipeIngredientService) *RecipeIngredientHandler {
	return &RecipeIngredientHandler{
		Logger:  logger,
		Service: service,
	}
}

func (h *RecipeIngredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	// r.Get("/", h.ViewIngredientList)
	r.Post("/create", h.Create)

	return r
}

func (h *RecipeIngredientHandler) Create(w http.ResponseWriter, r *http.Request) {
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
