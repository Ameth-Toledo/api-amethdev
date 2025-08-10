package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
	"errors"
)

type GetDonacionesByUsuario struct {
	db domain.IDonacion
}

func NewGetDonacionesByUsuario(db domain.IDonacion) *GetDonacionesByUsuario {
	return &GetDonacionesByUsuario{db: db}
}

func (gdbu *GetDonacionesByUsuario) Execute(usuarioID int) ([]entities.Donacion, error) {
	if usuarioID <= 0 {
		return nil, errors.New("usuario_id inválido")
	}

	return gdbu.db.GetByUsuarioID(usuarioID)
}
