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

	rh := handlers.NewRecipeHandler(log, services.NewRecipeService(log, ruow))
	hh := handlers.NewHomeHandler(log, services.NewRecipeService(log, ruow))
	ih := handlers.NewIngredientHandler(log, services.NewIngredientService(log, ids))
	qh := handlers.NewQuantityUnitHandler(log, services.NewQuantityUnitService(log, qds))
	ch := handlers.NewCategoryHandler(log, services.NewCategoryService(log, cds))

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
