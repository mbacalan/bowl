package models

import (
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
