package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
)

type GetComentariosByModulo struct {
	db domain.IComentario
}

func NewGetComentariosByModulo(db domain.IComentario) *GetComentariosByModulo {
	return &GetComentariosByModulo{db: db}
}

func (gcbm *GetComentariosByModulo) Execute(moduloId int) ([]entities.Comentario, error) {
	if moduloId <= 0 {
		return nil, errors.New("ID del módulo inválido")
	}
	return gcbm.db.GetByModuloId(moduloId)
}
