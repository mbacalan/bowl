package models

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	Router   *chi.Mux
	Database *gorm.DB
	Logger   *slog.Logger
	Repos    *Repositories
	Services *Services
	Handlers *Handlers
}
