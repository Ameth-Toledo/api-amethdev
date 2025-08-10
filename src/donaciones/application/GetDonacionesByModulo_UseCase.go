package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
	"errors"
)

type GetDonacionesByModulo struct {
	db domain.IDonacion
}

func NewGetDonacionesByModulo(db domain.IDonacion) *GetDonacionesByModulo {
	return &GetDonacionesByModulo{db: db}
}

func (gdbm *GetDonacionesByModulo) Execute(moduloID int) ([]entities.Donacion, error) {
	if moduloID <= 0 {
		return nil, errors.New("modulo_id inválido")
	}

	return gdbm.db.GetByModuloID(moduloID)
}
