package components

import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type Settings struct {
	IsAdmin bool
}

type contextKey string

const SessionKey contextKey = "bowl-session"

func GetSettings(r *http.Request) Settings {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, "bowl-session")
	// TODO: this can fail for reasons, add error handling
	isAdmin := session.Values["IsAdmin"].(bool)

	return Settings{IsAdmin: isAdmin}
}

func GetSessionValue(ctx context.Context, key string) any {
	if session, ok := ctx.Value(SessionKey).(*sessions.Session); ok {
		return session.Values[key]
	}

	return nil
}
