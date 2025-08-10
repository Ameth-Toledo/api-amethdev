package application

import (
	"AmethToledo/src/likes/domain"
	"errors"
	"time"
)

type GetLikeStats struct {
	db domain.ILike
}

func NewGetLikeStats(db domain.ILike) *GetLikeStats {
	return &GetLikeStats{db: db}
}

func (gls *GetLikeStats) Execute(moduloID int, startDate, endDate time.Time) (*domain.LikeStats, error) {
	if moduloID <= 0 {
		return nil, errors.New("el id del módulo es obligatorio y debe ser válido")
	}

	if startDate.After(endDate) {
		return nil, errors.New("la fecha de inicio debe ser anterior a la fecha de fin")
	}

	// Validar que el rango no sea mayor a 1 año
	if endDate.Sub(startDate) > 365*24*time.Hour {
		return nil, errors.New("el rango de fechas no puede ser mayor a 1 año")
	}

	return gls.db.GetLikeStatsByDateRange(moduloID, startDate, endDate)
}
