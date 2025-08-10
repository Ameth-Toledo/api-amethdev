package application

import (
	"AmethToledo/src/likes/domain"
	"AmethToledo/src/likes/domain/entities"
	"errors"
)

type GetLikesByModulo struct {
	db domain.ILike
}

func NewGetLikesByModulo(db domain.ILike) *GetLikesByModulo {
	return &GetLikesByModulo{db: db}
}

func (glbm *GetLikesByModulo) Execute(moduloID int) ([]entities.Like, error) {
	if moduloID <= 0 {
		return nil, errors.New("el id del módulo es obligatorio y debe ser válido")
	}

	return glbm.db.GetByModulo(moduloID)
}
