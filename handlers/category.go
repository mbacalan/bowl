package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/services"
)

type CategoryHandler struct {
	Logger  *slog.Logger
	Service services.CategoryService
}

func NewCategoryHandler(log *slog.Logger, service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		Logger:  log,
		Service: service,
	}
}

func (h *CategoryHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewAll)
	r.Get("/{id}", h.View)

	return r
}

func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.ViewAll(w, r)
}

func (h *CategoryHandler) View(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	category, err := h.Service.Get(id)
	if err != nil {
		h.Logger.Error("Error getting category", err)
	}

	pages.Category(category).Render(r.Context(), w)
}

func (h *CategoryHandler) ViewAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.Categories(categories).Render(r.Context(), w)
}
