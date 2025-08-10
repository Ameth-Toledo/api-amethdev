package application

import (
	"AmethToledo/src/comentarios/domain"
	"AmethToledo/src/comentarios/domain/entities"
)

type GetAllComentarios struct {
	db domain.IComentario
}

func NewGetAllComentarios(db domain.IComentario) *GetAllComentarios {
	return &GetAllComentarios{db: db}
}

func (gac *GetAllComentarios) Execute() ([]entities.Comentario, error) {
	return gac.db.GetAll()
}
