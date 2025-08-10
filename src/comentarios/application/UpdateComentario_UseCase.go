package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
)

type UpdateComentario struct {
	db domain.IComentario
}

func NewUpdateComentario(db domain.IComentario) *UpdateComentario {
	return &UpdateComentario{db: db}
}

func (uc *UpdateComentario) Execute(comentario entities.Comentario) error {
	if comentario.ID <= 0 {
		return errors.New("ID inválido")
	}
	if comentario.Texto == "" {
		return errors.New("el texto del comentario es obligatorio")
	}

	existingComentario, err := uc.db.GetById(comentario.ID)
	if err != nil {
		return err
	}
	if existingComentario == nil {
		return errors.New("comentario no encontrado")
	}

	return uc.db.Update(comentario)
}
