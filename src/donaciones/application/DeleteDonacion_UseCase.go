package application

import (
	"AmethToledo/src/donaciones/domain"
	"errors"
)

type DeleteDonacion struct {
	db domain.IDonacion
}

func NewDeleteDonacion(db domain.IDonacion) *DeleteDonacion {
	return &DeleteDonacion{db: db}
}

func (dd *DeleteDonacion) Execute(id int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	// Verificar que la donación existe
	existing, err := dd.db.GetById(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("donación no encontrada")
	}

	return dd.db.Delete(id)
}
