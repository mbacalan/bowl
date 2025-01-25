package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

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
	r.Put("/ingredients/create", h.CreateIngredient)
	r.Delete("/ingredients/{id}", h.DeleteIngredient)

	r.Get("/quantity-units", h.ViewQuantityUnits)
	r.Put("/quantity-units/create", h.CreateQuantityUnit)
	r.Delete("/quantity-units/{id}", h.DeleteQuantityUnit)

	return r
}

func (h *AdminHandler) View(w http.ResponseWriter, r *http.Request) {
	admin.Admin().Render(r.Context(), w)
}

func (h *AdminHandler) ViewIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetIngredients()
	if err != nil {
		h.Logger.Error("Error listing ingredients", "error", err)
	}

	admin.IngredientListPage(ingredients).Render(r.Context(), w)
}

func (h *AdminHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ingredient := r.Form["ingredient"]
	_, err := h.Service.CreateIngredient(ingredient[0])
	if err != nil {
		h.Logger.Error("Error creating ingredient", "error", err)
	}

	h.ViewIngredients(w, r)
}

func (h *AdminHandler) DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	h.Service.DeleteIngredient(uint(id))

	h.ViewIngredients(w, r)
}

func (h *AdminHandler) ViewQuantityUnits(w http.ResponseWriter, r *http.Request) {
	units, err := h.Service.GetQuantityUnits()

	if err != nil {
		h.Logger.Error("Error listing quantity units", "error", err)
		return
	}

	admin.QuantityUnitListPage(units).Render(r.Context(), w)
}

func (h *AdminHandler) CreateQuantityUnit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	quantityUnit := r.Form["quantity-unit"]
	_, err := h.Service.CreateQuantityUnit(quantityUnit[0])
	if err != nil {
		h.Logger.Error("Error creating quantity unit", "error", err)
	}

	h.ViewQuantityUnits(w, r)
}

func (h *AdminHandler) DeleteQuantityUnit(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	h.Service.DeleteQuantityUnit(uint(id))

	h.ViewQuantityUnits(w, r)
}
