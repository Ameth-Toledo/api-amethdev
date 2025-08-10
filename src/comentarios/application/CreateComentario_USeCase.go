package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
	"time"
)

type CreateComentario struct {
	db domain.IComentario
}

func NewCreateComentario(db domain.IComentario) *CreateComentario {
	return &CreateComentario{db: db}
}

func (cc *CreateComentario) Execute(comentario entities.Comentario) (entities.Comentario, error) {
	if comentario.ModuloID == 0 {
		return entities.Comentario{}, errors.New("el ID del módulo es obligatorio")
	}
	if comentario.UsuarioID == 0 {
		return entities.Comentario{}, errors.New("el ID del usuario es obligatorio")
	}
	if comentario.Texto == "" {
		return entities.Comentario{}, errors.New("el texto del comentario es obligatorio")
	}

	comentario.Fecha = time.Now()

	return cc.db.Save(comentario)
}
