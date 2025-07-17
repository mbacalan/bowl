package internal

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
)

type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	ErrorMessage   string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrUnathorized() render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessage:   "Unauthorized",
	}
}

func Authenticated(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "bowl-session")

			if session.IsNew {
				render.Render(w, r, ErrUnathorized())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func IsAdmin(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "bowl-session")

			if session.IsNew {
				render.Render(w, r, ErrUnathorized())
				return
			}

			isAdmin := session.Values["IsAdmin"].(bool)

			if !isAdmin {
				render.Render(w, r, ErrUnathorized())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
