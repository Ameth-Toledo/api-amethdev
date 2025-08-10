package application

import (
	"AmethToledo/src/contenidos/domain"
	"AmethToledo/src/contenidos/domain/entities"
)

type GetAllContenidos struct {
	db domain.IContenido
}

func NewGetAllContenidos(db domain.IContenido) *GetAllContenidos {
	return &GetAllContenidos{db: db}
}

func (gac *GetAllContenidos) Execute() ([]entities.Contenido, error) {
	return gac.db.GetAll()
}
