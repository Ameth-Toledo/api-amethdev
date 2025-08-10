package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
	"errors"
)

type GetDonacionById struct {
	db domain.IDonacion
}

func NewGetDonacionById(db domain.IDonacion) *GetDonacionById {
	return &GetDonacionById{db: db}
}

func (gdb *GetDonacionById) Execute(id int) (*entities.Donacion, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	donacion, err := gdb.db.GetById(id)
	if err != nil {
		return nil, err
	}

	if donacion == nil {
		return nil, errors.New("donación no encontrada")
	}

	return donacion, nil
}
