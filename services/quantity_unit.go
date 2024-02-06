package services

import (
	"log/slog"

	"github.com/mbacalan/bowl/db"
)

type QuantityUnitService struct {
	Log   *slog.Logger
	Store *db.QuantityUnitStore
}

func NewQuantityUnitService(log *slog.Logger, rs *db.QuantityUnitStore) QuantityUnitService {
	return QuantityUnitService{Log: log, Store: rs}
}

func (s *QuantityUnitService) GetAll() (units []db.QuantityUnit, error error) {
	result, err := s.Store.GetAll()

	if err != nil {
		s.Log.Error("Error getting all quantity units", err)
		return result, err
	}

	return result, nil
}

func (s *QuantityUnitService) Create(i string) (unit db.QuantityUnit, error error) {
	result, err := s.Store.GetOrCreate(i)

	if err != nil {
		s.Log.Error("Error creating unit", err)
		return result, err
	}

	return result, nil
}
