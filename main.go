package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mbacalan/bowl/assets"
	"github.com/mbacalan/bowl/components/pages"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", templ.Handler(pages.Home()).ServeHTTP)
	r.Get("/recipe", templ.Handler(pages.Recipe()).ServeHTTP)

	assets.Mount(r)

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.TimeoutHandler(r, 30*time.Second, "request timed out"),
	}

	fmt.Printf("Listening on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to start the server", err)
		os.Exit(1)
	}
}
