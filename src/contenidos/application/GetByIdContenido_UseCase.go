package application

import (
	"AmethToledo/src/contenidos/domain"
	"AmethToledo/src/contenidos/domain/entities"
	"errors"
)

type GetContenidoById struct {
	db domain.IContenido
}

func NewGetContenidoById(db domain.IContenido) *GetContenidoById {
	return &GetContenidoById{db: db}
}

func (gcb *GetContenidoById) Execute(id int) (*entities.Contenido, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	return gcb.db.GetById(id)
}
