package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/recipes"
	"github.com/mbacalan/bowl/models"
)

type recipeHandler struct {
	*models.RecipeHandler
}

func NewRecipeHandler(logger *slog.Logger, service models.RecipeService) *recipeHandler {
	return &recipeHandler{
		RecipeHandler: &models.RecipeHandler{
			Logger:  logger,
			Service: service,
		},
	}
}

func (h *recipeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)
	r.Get("/{id}", h.View)
	r.Get("/{id}/edit", h.Edit)
	r.Patch("/{id}", h.Update)
	r.Get("/create", h.Create)
	r.Post("/create", h.Create)

	return r
}

func (h *recipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		recipes.CreateRecipe().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	name := r.Form.Get("name")
	prepDuration, _ := strconv.ParseUint(r.Form.Get("prep-duration"), 10, 32)
	cookDuration, _ := strconv.ParseUint(r.Form.Get("cook-duration"), 10, 32)
	steps := r.Form["step"]
	categories := r.Form.Get("categories")
	ingredients := r.Form["ingredient"]
	quantities := r.Form["quantity"]
	quantityUnits := r.Form["quantity-unit"]

	recipe, err := h.Service.Create(models.RecipeData{
		Name:          cases.Title(language.English).String(name),
		PrepDuration:  uint(prepDuration),
		CookDuration:  uint(cookDuration),
		Steps:         steps,
		Categories:    strings.Split(categories, ", "),
		Ingredients:   ingredients,
		Quantities:    quantities,
		QuantityUnits: quantityUnits,
	})

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(recipe.ID), 10))
	recipeDetail, _ := h.Service.Get(int(recipe.ID))
	recipes.RecipeDetailPage(recipeDetail).Render(r.Context(), w)
}

func (h *recipeHandler) View(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	recipe, err := h.Service.Get(id)

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	recipes.RecipeDetailPage(recipe).Render(r.Context(), w)
}

func (h *recipeHandler) Edit(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	recipe, err := h.Service.Get(id)

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	recipes.EditRecipe(recipe).Render(r.Context(), w)
}

func (h *recipeHandler) Update(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	r.ParseForm()

	name := r.Form.Get("name")
	prepDuration, _ := strconv.ParseUint(r.Form.Get("prep-duration"), 10, 32)
	cookDuration, _ := strconv.ParseUint(r.Form.Get("cook-duration"), 10, 32)
	steps := r.Form["step"]
	categories := r.Form.Get("categories")
	ingredients := r.Form["ingredient"]
	quantities := r.Form["quantity"]
	quantityUnits := r.Form["quantity-unit"]

	data := models.RecipeData{
		Name:          name,
		PrepDuration:  uint(prepDuration),
		CookDuration:  uint(cookDuration),
		Steps:         steps,
		Categories:    strings.Split(categories, ", "),
		Ingredients:   ingredients,
		Quantities:    quantities,
		QuantityUnits: quantityUnits,
	}

	w.Header().Set("HX-Push-URL", "/recipes/"+strconv.FormatUint(uint64(id), 10))
	recipeDetail, _ := h.Service.Update(id, data)
	recipes.RecipeDetailPage(recipeDetail).Render(r.Context(), w)
}

func (h *recipeHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	rs, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing recipes", err)
	}

	recipes.RecipeListPage(rs).Render(r.Context(), w)
}
