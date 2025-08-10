package application

import (
	"AmethToledo/src/likes/domain"
	"errors"
)

type GetMostLikedModulos struct {
	db domain.ILike
}

func NewGetMostLikedModulos(db domain.ILike) *GetMostLikedModulos {
	return &GetMostLikedModulos{db: db}
}

func (gmlm *GetMostLikedModulos) Execute(limit int) ([]domain.ModuloWithLikes, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		return nil, errors.New("el límite no puede ser mayor a 100")
	}

	return gmlm.db.GetMostLikedModulos(limit)
}
