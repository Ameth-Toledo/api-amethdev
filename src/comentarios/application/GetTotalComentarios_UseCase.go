package application

import "AmethToledo/src/comentarios/domain"

type GetTotalComentarios struct {
	db domain.IComentario
}

func NewGetTotalComentarios(db domain.IComentario) *GetTotalComentarios {
	return &GetTotalComentarios{db: db}
}

func (gtc *GetTotalComentarios) Execute() (int, error) {
	return gtc.db.GetTotal()
}
