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
	"github.com/mbacalan/bowl/db"
	"github.com/mbacalan/bowl/handlers"
	"github.com/mbacalan/bowl/services"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	database := db.NewConnection()
	rds := db.NewRecipeStore(database, "recipes")
	ids := db.NewIngredientStore(database, "ingredients")

	rs := services.NewRecipeService(log, rds)
	rh := handlers.NewRecipeHandler(log, rs)
	is := services.NewIngredientService(log, ids)
	ih := handlers.NewIngredientHandler(log, is)

	recent, _ := rs.GetRecent(5)
	r.Get("/", templ.Handler(pages.Home(recent)).ServeHTTP)
	r.Mount("/assets", assets.Routes())
	r.Mount("/recipes", rh.Routes())
	r.Mount("/ingredients", ih.Routes())

	server := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	fmt.Printf("Listening on %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Error("Failed to start the server", err)
		os.Exit(1)
	}
}
