package models

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

type Handlers struct {
	AuthHandler         AuthHandlerInterface
	AdminHandler        AdminHandlerInterface
	HomeHandler         HomeHandlerInterface
	RecipeHandler       RecipeHandlerInterface
	QuantityUnitHandler QuantityUnitHandlerInterface
	CategoryHandler     CategoryHandlerInterface
}

type AuthHandlerInterface interface {
	Routes() chi.Router
	GetStore() *sessions.CookieStore
}

type AdminHandlerInterface interface {
	Routes() chi.Router
}

type CategoryHandlerInterface interface {
	Routes() chi.Router
}

type HomeHandlerInterface interface {
	Routes() chi.Router
}

type QuantityUnitHandlerInterface interface {
	Routes() chi.Router
}

type RecipeHandlerInterface interface {
	Routes() chi.Router
}
