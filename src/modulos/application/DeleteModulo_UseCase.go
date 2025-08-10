package application

import (
	"AmethToledo/src/modulos/domain"
	"errors"
)

type DeleteModulo struct {
	db domain.IModulo
}

func NewDeleteModulo(db domain.IModulo) *DeleteModulo {
	return &DeleteModulo{db: db}
}

func (dm *DeleteModulo) Execute(id int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	return dm.db.Delete(id)
}
