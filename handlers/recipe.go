package handlers

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/db"
	"github.com/mbacalan/bowl/services"
)

type RecipeHandler struct {
	Log           *slog.Logger
	RecipeService services.Service
}

func New(service services.Service) RecipeHandler {
	return RecipeHandler{RecipeService: service}
}

func (h *RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Create(w, r)
		return
	}

	h.ViewList(w, r)
}

func Mount(r chi.Router, h RecipeHandler) {
	r.Route("/recipes", func(r chi.Router) {
		r.Get("/", h.ViewList)
		r.Get("/{id}", h.ViewRecipe)
		r.Get("/create", h.Create)
		r.Post("/create", h.Create)
	})
}

func (h *RecipeHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateRecipe().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	recipe, err := db.CreateRecipe(db.Recipe{Name: r.Form.Get("name")})

	if err != nil {
		log.Fatal(err)
		return
	}

	pages.RecipeDetail(recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) ViewRecipe(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	recipe, err := h.RecipeService.Get(id)

	if err != nil {
		h.Log.Error("", err)
		return
	}

	pages.RecipeDetail(recipe).Render(r.Context(), w)
}

func (h *RecipeHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.RecipeService.GetAll()
	if err != nil {
		h.Log.Error("Error listing recipes", err)
	}

	pages.RecipeList(recipes).Render(r.Context(), w)
}
