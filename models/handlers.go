package models

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	HomeHandler         HomeHandlerInterface
	AuthHandler         AuthHandlerInterface
	RecipeHandler       RecipeHandlerInterface
	IngredientHandler   IngredientHandlerInterface
	QuantityUnitHandler QuantityUnitHandlerInterface
	CategoryHandler     CategoryHandlerInterface
}

type AuthHandler struct {
	Logger  *slog.Logger
	Service AuthService
	Store   *sessions.CookieStore
}

type CategoryHandler struct {
	Logger  *slog.Logger
	Service CategoryService
}

type HomeHandler struct {
	Logger  *slog.Logger
	Service HomeService
}

type IngredientHandler struct {
	Logger  *slog.Logger
	Service IngredientService
}

type QuantityUnitHandler struct {
	Logger  *slog.Logger
	Service QuantityUnitService
}

type RecipeHandler struct {
	Logger  *slog.Logger
	Service RecipeService
}

type AuthHandlerInterface interface {
	Routes() chi.Router
	GetStore() *sessions.CookieStore
}

type CategoryHandlerInterface interface {
	Routes() chi.Router
}

type HomeHandlerInterface interface {
	Routes() chi.Router
}

type IngredientHandlerInterface interface {
	Routes() chi.Router
}

type QuantityUnitHandlerInterface interface {
	Routes() chi.Router
}

type RecipeHandlerInterface interface {
	Routes() chi.Router
}
