package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
)

type GetComentariosByUsuario struct {
	db domain.IComentario
}

func NewGetComentariosByUsuario(db domain.IComentario) *GetComentariosByUsuario {
	return &GetComentariosByUsuario{db: db}
}

func (gcbu *GetComentariosByUsuario) Execute(usuarioId int) ([]entities.Comentario, error) {
	if usuarioId <= 0 {
		return nil, errors.New("ID del usuario inválido")
	}
	return gcbu.db.GetByUsuarioId(usuarioId)
}
