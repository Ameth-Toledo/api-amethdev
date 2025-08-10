package application

import (
	"AmethToledo/src/modulos/domain"
	"AmethToledo/src/modulos/domain/entities"
)

type GetAllModulos struct {
	db domain.IModulo
}

func NewGetAllModulos(db domain.IModulo) *GetAllModulos {
	return &GetAllModulos{db: db}
}

func (gam *GetAllModulos) Execute() ([]entities.Modulo, error) {
	return gam.db.GetAll()
}
