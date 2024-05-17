package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type categoryHandler struct {
	*models.CategoryHandler
}

func NewCategoryHandler(logger *slog.Logger, service models.CategoryService) *categoryHandler {
	return &categoryHandler{
		CategoryHandler: &models.CategoryHandler{
			Logger:  logger,
			Service: service,
		},
	}
}

func (h *categoryHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)
	r.Get("/{id}", h.View)

	return r
}

func (h *categoryHandler) View(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	category, err := h.Service.Get(id)
	if err != nil {
		h.Logger.Error("Error getting category", err)
	}

	pages.Category(category).Render(r.Context(), w)
}

func (h *categoryHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.Categories(categories).Render(r.Context(), w)
}
