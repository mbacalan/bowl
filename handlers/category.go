package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type CategoryHandler struct {
	Logger  *slog.Logger
	Service CategoryService
}

type CategoryService interface {
	Get(id int) (models.Category, error)
	GetAll() ([]models.Category, error)
}

func NewCategoryHandler(logger *slog.Logger, service CategoryService) *CategoryHandler {
	return &CategoryHandler{
		Logger:  logger,
		Service: service,
	}
}

func (h *CategoryHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewList)
	r.Get("/{id}", h.View)

	return r
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

func (h *CategoryHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.GetAll()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.Categories(categories).Render(r.Context(), w)
}
