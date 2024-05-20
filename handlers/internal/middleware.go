package internal

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func Authenticated(store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, "bowl-session")

			if session.IsNew {
				http.Redirect(w, r, "/auth", http.StatusFound)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
