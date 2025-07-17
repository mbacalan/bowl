package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"github.com/mbacalan/bowl/models"
)

// Standard API response for success
type SuccessResponse struct {
	HTTPStatusCode int `json:"-"`
}

// Set 200 OK for rendering the response
func (s *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

// Return a new response with 200 OK, in case you need to access it before Render()
func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{
		HTTPStatusCode: http.StatusOK,
	}
}

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	ErrorMessage   string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage:   err.Error(),
	}
}

func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage:   err.Error(),
	}
}

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

	r.Post("/signup", h.Signup)
	r.Post("/login", h.Login)
	r.Post("/logout", h.Logout)

	return r
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := h.Service.Signup(username, password)

	if err != nil {
		h.Logger.Error("Error signing up", "error", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	session := h.createSession(w, r, &user)
	ctx := context.WithValue(r.Context(), "bowl-session", session)
	if err := render.Render(w, r.WithContext(ctx), NewSuccessResponse()); err != nil {
		render.Render(w, r, ErrInternal(err))
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := h.Service.Login(username, password)

	if err != nil {
		h.Logger.Error("Error logging in", "error", err)
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	session := h.createSession(w, r, &user)
	ctx := context.WithValue(r.Context(), "bowl-session", session)
	if err := render.Render(w, r.WithContext(ctx), NewSuccessResponse()); err != nil {
		render.Render(w, r, ErrInternal(err))
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := h.Store.Get(r, "bowl-session")
	session.Options.MaxAge = -1

	err := session.Save(r, w)
	if err != nil {
		h.Logger.Error("Error logging out", "error", err)
		render.Render(w, r, ErrInternal(err))
		return
	}

	render.Render(w, r, NewSuccessResponse())
}

func (h *AuthHandler) GetStore() *sessions.CookieStore {
	return h.Store
}

func (h *AuthHandler) createSession(w http.ResponseWriter, r *http.Request, user *models.User) *sessions.Session {
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
		render.Render(w, r, ErrInternal(err))
	}

	return session
}
