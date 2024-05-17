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
	Service AuthService
	Store   *sessions.CookieStore
}

type AuthService interface {
	Signup(name string, password string) (models.User, error)
	Login(name string, password string) (models.User, error)
}

func NewAuthHandler(logger *slog.Logger, service AuthService) *AuthHandler {
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

	_, err := h.Service.Signup(username, password)

	if err != nil {
		h.Logger.Error("Error signing up in", err)
	}

	h.createSession(w, r)
	w.Header().Set("HX-Push-URL", "/")
	pages.Home([]models.Recipe{}).Render(r.Context(), w)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err := h.Service.Login(username, password)

	if err != nil {
		h.Logger.Error("Error logging in", err)
	}

	h.createSession(w, r)
	w.Header().Set("HX-Push-URL", "/")
	pages.Home([]models.Recipe{}).Render(r.Context(), w)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := h.Store.Get(r, "bowl-session")
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/auth", http.StatusFound)
}

func (h *AuthHandler) createSession(w http.ResponseWriter, r *http.Request) {
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

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
