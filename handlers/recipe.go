package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components/recipes"
	"github.com/mbacalan/bowl/models"
)

type RecipeHandler struct {
	Logger  *slog.Logger
	Service models.RecipeService
	Store   *sessions.CookieStore
}

func NewRecipeHandler(logger *slog.Logger, service models.RecipeService, store *sessions.CookieStore) *RecipeHandler {
	return &RecipeHandler{
		Logger:  logger,
		Service: service,
		Store:   store,
	}
}

func (h *RecipeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)
	r.Get("/{id}", h.View)
	r.Get("/{id}/edit", h.Edit)
	r.Patch("/{id}", h.Update)
	r.Get("/create", h.Create)
	r.Post("/create", h.Create)

	return r
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)

	recipe, err := h.Service.Create(models.RecipeData{
		Name:          cases.Title(language.English).String(name),
		PrepDuration:  uint(prepDuration),
		CookDuration:  uint(cookDuration),
		Steps:         steps,
		Categories:    strings.Split(categories, ", "),
		Ingredients:   ingredients,
		Quantities:    quantities,
		QuantityUnits: quantityUnits,
		UserID:        user,
	})

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(recipe.ID), 10))
	recipeDetail, _ := h.Service.Get(user, int(recipe.ID))
	recipes.RecipeDetailPage(recipeDetail).Render(r.Context(), w)
}

func (h *RecipeHandler) View(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)
	recipe, err := h.Service.Get(user, id)

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	recipes.RecipeDetailPage(recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) Edit(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)

	recipe, err := h.Service.Get(user, id)

	if err != nil {
		h.Logger.Error("", err)
		return
	}

	recipes.EditRecipe(recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *RecipeHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)

	rs, err := h.Service.GetAll(user)
	if err != nil {
		h.Logger.Error("Error listing recipes", err)
	}

	recipes.RecipeListPage(rs).Render(r.Context(), w)
}
