package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type HomeHandler struct {
	Logger  *slog.Logger
	Service models.HomeService
	Store   *sessions.CookieStore
}

func NewHomeHandler(logger *slog.Logger, service models.HomeService, store *sessions.CookieStore) *HomeHandler {
	return &HomeHandler{
		Logger:  logger,
		Service: service,
		Store:   store,
	}
}

func (h *HomeHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.View)

	return r
}

func (h *HomeHandler) View(w http.ResponseWriter, r *http.Request) {
	session, err := h.Store.Get(r, "bowl-session")
	if err != nil {
		h.Logger.Error("Error getting user session", "error", err)
		return
	}

	user := session.Values["UserID"].(uint)
	recipes, err := h.Service.GetRecent(user, 10)

	if err != nil {
		h.Logger.Error("Error viewing home", "error", err)
	}

	pages.Home(recipes).Render(r.Context(), w)
}
