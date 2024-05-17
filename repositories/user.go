package repositories

import (
	"github.com/mbacalan/bowl/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	tableName string
}

func NewUserRepository(db *gorm.DB, tableName string) *UserRepository {
	repository := &UserRepository{
		tableName: tableName,
		db:        db,
	}

	return repository
}

func (s UserRepository) Create(name string, password []byte) (models.User, error) {
	entry := models.User{Name: name, Password: password}

	error := s.db.Create(&entry).Error

	return entry, error
}

func (s UserRepository) Get(name string) (models.User, error) {
	var entry models.User

	error := s.db.First(&entry, "name = ?", name).Error

	return entry, error
}
