package application

import (
	"AmethToledo/src/donaciones/domain"
	"errors"
)

type GetStatsDonaciones struct {
	db domain.IDonacion
}

func NewGetStatsDonaciones(db domain.IDonacion) *GetStatsDonaciones {
	return &GetStatsDonaciones{db: db}
}

type StatsResponse struct {
	TotalByUsuario float64 `json:"total_by_usuario"`
	TotalByModulo  float64 `json:"total_by_modulo"`
}

func (gsd *GetStatsDonaciones) ExecuteByUsuario(usuarioID int) (float64, error) {
	if usuarioID <= 0 {
		return 0, errors.New("usuario_id inválido")
	}

	return gsd.db.GetTotalByUsuario(usuarioID)
}

func (gsd *GetStatsDonaciones) ExecuteByModulo(moduloID int) (float64, error) {
	if moduloID <= 0 {
		return 0, errors.New("modulo_id inválido")
	}

	return gsd.db.GetTotalByModulo(moduloID)
}
