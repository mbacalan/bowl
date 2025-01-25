package handlers

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/components/pages"
	"github.com/mbacalan/bowl/models"
)

type AuthHandler struct {
	Logger  *slog.Logger
	Service models.AuthService
	Store   *sessions.CookieStore
}

func NewAuthHandler(logger *slog.Logger, service models.AuthService) *AuthHandler {
	return &AuthHandler{
		Logger:  logger,
		Service: service,
		Store:   sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET"))),
	}
}

func (h *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.Auth)
	r.Post("/signup", h.Signup)
	r.Post("/login", h.Login)
	r.Get("/logout", h.Logout)

	return r
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	pages.Auth().Render(r.Context(), w)
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := h.Service.Signup(username, password)

	if err != nil {
		h.Logger.Error("Error signing up in", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	h.createSession(w, r, &user)
	w.Header().Set("HX-Push-URL", "/")
	pages.Home([]models.Recipe{}).Render(r.Context(), w)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := h.Service.Login(username, password)

	if err != nil {
		h.Logger.Error("Error logging in", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	h.createSession(w, r, &user)
	w.Header().Set("HX-Push-URL", "/")
	pages.Home([]models.Recipe{}).Render(r.Context(), w)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := h.Store.Get(r, "bowl-session")
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		h.Logger.Error("Error logging out", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
		return
	}

	http.Redirect(w, r, "/auth", http.StatusFound)
}

func (h *AuthHandler) GetStore() *sessions.CookieStore {
	return h.Store
}

func (h *AuthHandler) createSession(w http.ResponseWriter, r *http.Request, user *models.User) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := h.Store.Get(r, "bowl-session")
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 30,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	session.Values["UserID"] = user.ID
	session.Values["UserName"] = user.Name
	session.Values["IsAdmin"] = user.IsAdmin
	err := session.Save(r, w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		pages.Error(err.Error()).Render(r.Context(), w)
	}
}
