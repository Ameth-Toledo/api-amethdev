package application

import (
	"AmethToledo/src/modulos/domain"
	"AmethToledo/src/modulos/domain/entities"
	"errors"
)

type GetModuloById struct {
	db domain.IModulo
}

func NewGetModuloById(db domain.IModulo) *GetModuloById {
	return &GetModuloById{db: db}
}

func (gmb *GetModuloById) Execute(id int) (*entities.Modulo, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	return gmb.db.GetById(id)
}
