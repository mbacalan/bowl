package assets

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed *
var Assets embed.FS

func Routes() chi.Router {
	fs := http.FileServer(http.FS(Assets))
	r := chi.NewRouter()

	r.Handle("/*", http.StripPrefix("/assets/", fs))

	return r
}
