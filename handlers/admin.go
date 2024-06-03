package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components/admin"
	"github.com/mbacalan/bowl/models"
)

type AdminHandler struct {
	Logger  *slog.Logger
	Service models.AdminService
	Store   *sessions.CookieStore
}

func NewAdminHandler(logger *slog.Logger, service models.AdminService, store *sessions.CookieStore) *AdminHandler {
	return &AdminHandler{
		Logger:  logger,
		Service: service,
		Store:   store,
	}
}

func (h *AdminHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.View)
	r.Get("/ingredients", h.ViewIngredients)
	r.Post("/ingredients/create", h.CreateIngredient)

	return r
}

func (h *AdminHandler) View(w http.ResponseWriter, r *http.Request) {
	admin.Admin().Render(r.Context(), w)
}

func (h *AdminHandler) ViewIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetIngredients()
	if err != nil {
		h.Logger.Error("Error listing ingredients", err)
	}

	admin.IngredientListPage(ingredients).Render(r.Context(), w)
}

func (h *AdminHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ingredient := r.Form["ingredient"]
	_, err := h.Service.CreateIngredient(ingredient[0])
	if err != nil {
		h.Logger.Error("Error creating ingredient", err)
	}

	http.Redirect(w, r, "/admin/ingredients", http.StatusFound)
}
