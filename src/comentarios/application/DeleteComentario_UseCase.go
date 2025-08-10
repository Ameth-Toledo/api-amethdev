package application

import (
	"AmethToledo/src/comentarios/domain"
	"errors"
)

type DeleteComentario struct {
	db domain.IComentario
}

func NewDeleteComentario(db domain.IComentario) *DeleteComentario {
	return &DeleteComentario{db: db}
}

func (dc *DeleteComentario) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	existingComentario, err := dc.db.GetById(id)
	if err != nil {
		return err
	}
	if existingComentario == nil {
		return errors.New("comentario no encontrado")
	}

	return dc.db.Delete(id)
}
