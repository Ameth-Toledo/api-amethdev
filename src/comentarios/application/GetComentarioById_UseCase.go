package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
)

type GetComentarioById struct {
	db domain.IComentario
}

func NewGetComentarioById(db domain.IComentario) *GetComentarioById {
	return &GetComentarioById{db: db}
}

func (gcbi *GetComentarioById) Execute(id int) (*entities.Comentario, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}
	return gcbi.db.GetById(id)
}
