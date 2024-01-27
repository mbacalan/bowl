package assets

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed *
var Assets embed.FS

// func Mount(r chi.Router) {
// 	r.Route("/assets", func(r chi.Router) {
// 		r.Use(func(next http.Handler) http.Handler {
// 			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				next.ServeHTTP(w, r)
// 			})
// 		})

// 		r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.FS(Assets))))
// 	})
// }

func Mount(r chi.Router) {
	fs := http.FileServer(http.FS(Assets))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))
}
