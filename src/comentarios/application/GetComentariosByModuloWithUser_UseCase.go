package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
	"errors"
)

type GetComentariosByModuloWithUser struct {
	db domain.IComentario
}

func NewGetComentariosByModuloWithUser(db domain.IComentario) *GetComentariosByModuloWithUser {
	return &GetComentariosByModuloWithUser{db: db}
}

func (gcbmwu *GetComentariosByModuloWithUser) Execute(moduloId int) ([]entities.ComentarioConUsuario, error) {
	if moduloId <= 0 {
		return nil, errors.New("ID del módulo inválido")
	}
	return gcbmwu.db.GetByModuloIdWithUser(moduloId)
}
