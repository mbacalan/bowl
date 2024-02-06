package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services"
)

type RecipeHandler struct {
	Log           *slog.Logger
	RecipeService services.RecipeService
}

func NewRecipeHandler(log *slog.Logger, service services.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		Log:           log,
		RecipeService: service,
	}
}

func (h *RecipeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)
	r.Get("/{id}", h.ViewRecipe)
	r.Get("/create", h.Create)
	r.Post("/create", h.Create)

	return r
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Create(w, r)
		return
	}

	h.ViewList(w, r)
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateRecipe().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	recipe, err := h.RecipeService.Create(db.Recipe{
		Name: r.Form.Get("name"),
	})

	ingredients := r.Form["ingredient"]
	quantities := r.Form["quantity"]
	quantityUnits := r.Form["quantity-unit"]

	for i := range ingredients {
		ingredient, _ := h.RecipeService.IngredientRepository.GetOrCreate(ingredients[i])
		unit, _ := h.RecipeService.QuantityUnitRepository.GetOrCreate(quantityUnits[i])

		h.RecipeService.RecipeIngredientRepository.Create(recipe.ID, ingredient.ID, unit.ID, quantities[i])
	}

	if err != nil {
		h.Log.Error("", err)
		return
	}

	w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(recipe.ID), 10))
	recipeDetail, _ := h.RecipeService.Get(int(recipe.ID))
	pages.RecipeDetailPage(recipeDetail).Render(r.Context(), w)
}

func (h *RecipeHandler) ViewRecipe(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	recipe, err := h.RecipeService.Get(id)

	if err != nil {
		h.Log.Error("", err)
		return
	}

	pages.RecipeDetailPage(recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.RecipeService.GetAll()
	if err != nil {
		h.Log.Error("Error listing recipes", err)
	}

	pages.RecipeListPage(recipes).Render(r.Context(), w)
}
