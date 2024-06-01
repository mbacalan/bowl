package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type CategoryHandler struct {
	Logger  *slog.Logger
	Service models.CategoryService
	Store   *sessions.CookieStore
}

func NewCategoryHandler(logger *slog.Logger, service models.CategoryService, store *sessions.CookieStore) *CategoryHandler {
	return &CategoryHandler{
		Logger:  logger,
		Service: service,
		Store:   store,
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
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)
	category, err := h.Service.Get(user, id)
	if err != nil {
		h.Logger.Error("Error getting category", err)
	}

	pages.Category(category).Render(r.Context(), w)
}

func (h *CategoryHandler) ViewList(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", err)
		return
	}

	user := session.Values["UserID"].(uint)
	categories, err := h.Service.GetAll(user)
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	pages.Categories(categories).Render(r.Context(), w)
}
