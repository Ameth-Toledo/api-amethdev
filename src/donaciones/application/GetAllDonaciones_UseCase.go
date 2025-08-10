package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
)

type GetAllDonaciones struct {
	db domain.IDonacion
}

func NewGetAllDonaciones(db domain.IDonacion) *GetAllDonaciones {
	return &GetAllDonaciones{db: db}
}

func (gad *GetAllDonaciones) Execute() ([]entities.Donacion, error) {
	return gad.db.GetAll()
}
