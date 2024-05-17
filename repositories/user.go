package repositories

import (
	"gorm.io/gorm"

	"github.com/mbacalan/bowl/models"
)

type UserRepository struct {
	*models.Repository
}

func NewUserRepository(db *gorm.DB, tableName string) *UserRepository {
	return &UserRepository{
		Repository: &models.Repository{
			DB:        db,
			TableName: tableName,
		},
	}
}

func (s *UserRepository) Create(name string, password []byte) (models.User, error) {
	entry := models.User{Name: name, Password: password}

	error := s.DB.Create(&entry).Error

	return entry, error
}

func (s *UserRepository) Get(name string) (models.User, error) {
	var entry models.User

	error := s.DB.First(&entry, "name = ?", name).Error

	return entry, error
}
