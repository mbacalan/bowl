package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components"
	"github.com/mbacalan/bowl/components/pages"
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

func (h *RecipeHandler) Settings(r *http.Request) components.Settings {
	settings := components.GetSettings(r)
	return components.Settings{IsAdmin: settings.IsAdmin}
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
		recipes.CreateRecipe(h.Settings(r)).Render(r.Context(), w)
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
		h.Logger.Error("Error getting user session", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
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
		h.Logger.Error("Error creating recipe", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(recipe.ID), 10))
	recipeDetail, _ := h.Service.Get(user, int(recipe.ID))
	recipes.Recipe(recipeDetail).Render(r.Context(), w)
}

func (h *RecipeHandler) View(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	user := session.Values["UserID"].(uint)
	recipe, err := h.Service.Get(user, id)
	if err != nil {
		h.Logger.Error("Error getting recipe", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	recipes.RecipeDetailPage(h.Settings(r), recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) Edit(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	user := session.Values["UserID"].(uint)

	recipe, err := h.Service.Get(user, id)
	if err != nil {
		h.Logger.Error("Error getting recipe", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	recipes.EditRecipe(h.Settings(r), recipe).Render(r.Context(), w)
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

	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	user := session.Values["UserID"].(uint)

	data := models.RecipeData{
		Name:          name,
		PrepDuration:  uint(prepDuration),
		CookDuration:  uint(cookDuration),
		Steps:         steps,
		Categories:    strings.Split(categories, ", "),
		Ingredients:   ingredients,
		Quantities:    quantities,
		QuantityUnits: quantityUnits,
		UserID:        user,
	}

	recipeDetail, err := h.Service.Update(id, data)

	if err != nil {
		h.Logger.Error("Error editing recipe", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Push-URL", "/recipes/"+strconv.FormatUint(uint64(id), 10))
	recipes.Recipe(recipeDetail).Render(r.Context(), w)
}

type RecipeResponse struct {
	models.Recipe
}

func (rd *RecipeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewRecipeResponse(recipe models.Recipe) *RecipeResponse {
	resp := &RecipeResponse{Recipe: recipe}
	return resp
}

func NewRecipeListResponse(recipes []models.Recipe) []render.Renderer {
	list := []render.Renderer{}

	for _, recipe := range recipes {
		list = append(list, NewRecipeResponse(recipe))
	}

	return list
}

func (h *RecipeHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "bowl-session")
	// TODO: user chi/render like in handlers/auth.go
	if err != nil {
		h.Logger.Error("Error getting user session", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	user := session.Values["UserID"].(uint)

	rs, err := h.Service.GetAll(user)
	if err != nil {
		h.Logger.Error("Error getting all recipes", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	render.RenderList(w, r, NewRecipeListResponse(rs))
}
