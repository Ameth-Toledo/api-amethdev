package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
)

type GetAllComentariosWithUser struct {
	db domain.IComentario
}

func NewGetAllComentariosWithUser(db domain.IComentario) *GetAllComentariosWithUser {
	return &GetAllComentariosWithUser{db: db}
}

func (gacwu *GetAllComentariosWithUser) Execute() ([]entities.ComentarioConUsuario, error) {
	return gacwu.db.GetAllWithUser()
}
