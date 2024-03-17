package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mbacalan/bowl/assets"
	"github.com/mbacalan/bowl/handlers"
	"github.com/mbacalan/bowl/repositories"
	"github.com/mbacalan/bowl/services"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	database := db.NewConnection()
	ids := db.NewIngredientRepository(database, "ingredients")
	qds := db.NewQuantityUnitRepository(database, "quantity_units")
	cds := db.NewCategoryRepository(database, "categories")
	ruow := services.NewRecipeUOW(database)

	rs := services.NewRecipeService(log, ruow)
	rh := handlers.NewRecipeHandler(log, rs)
	hh := handlers.NewHomeHandler(log, rs)
	is := services.NewIngredientService(log, ids)
	ih := handlers.NewIngredientHandler(log, is)
	qs := services.NewQuantityUnitService(log, qds)
	qh := handlers.NewQuantityUnitHandler(log, qs)
	cs := services.NewCategoryService(log, cds)
	ch := handlers.NewCategoryHandler(log, cs)

	r.Mount("/", hh.Routes())
	r.Mount("/assets", assets.Routes())
	r.Mount("/recipes", rh.Routes())
	r.Mount("/categories", ch.Routes())
	r.Mount("/ingredients", ih.Routes())
	r.Mount("/quantity-units", qh.Routes())

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
