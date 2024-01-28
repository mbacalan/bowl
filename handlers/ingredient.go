package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/db"
	"github.com/mbacalan/bowl/services"
)

type IngredientHandler struct {
	Log               *slog.Logger
	IngredientService services.IngredientServiceImpl
}

func NewIngredientHandler(log *slog.Logger, service services.IngredientServiceImpl) *IngredientHandler {
	return &IngredientHandler{
		Log:               log,
		IngredientService: service,
	}
}

func (h *IngredientHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.ViewIngredientList)
	r.Get("/{id}", h.ViewIngredient)
	r.Get("/create", h.CreateIngredient)
	r.Post("/create", h.CreateIngredient)

	return r
}

func (h *IngredientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.CreateIngredient(w, r)
		return
	}

	h.ViewIngredientList(w, r)
}

func (h *IngredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		pages.CreateIngredient().Render(r.Context(), w)
		return
	}

	r.ParseForm()

	ingredient, err := h.IngredientService.Create(db.Ingredient{Name: r.Form.Get("name")})

	if err != nil {
		h.Log.Error("", err)
		return
	}

	w.Header().Set("HX-Push-URL", strconv.FormatUint(uint64(ingredient.ID), 10))
	pages.IngredientDetailPage(ingredient).Render(r.Context(), w)
}

func (h *IngredientHandler) ViewIngredient(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(param)

	ingredient, err := h.IngredientService.Get(id)

	if err != nil {
		h.Log.Error("", err)
		return
	}

	pages.IngredientDetailPage(ingredient).Render(r.Context(), w)
}

func (h *IngredientHandler) ViewIngredientList(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.IngredientService.GetAll()
	if err != nil {
		h.Log.Error("Error listing ingredients", err)
	}

	pages.IngredientListPage(ingredients).Render(r.Context(), w)
}
