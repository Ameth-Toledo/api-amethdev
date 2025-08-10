package application

import "AmethToledo/src/donaciones/domain"

type GetTotalDonaciones struct {
	db domain.IDonacion
}

func NewGetTotalDonaciones(db domain.IDonacion) *GetTotalDonaciones {
	return &GetTotalDonaciones{db: db}
}

func (gtd *GetTotalDonaciones) Execute() (int, error) {
	return gtd.db.GetTotal()
}
