package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
	"errors"
	"time"
)

type CreateDonacion struct {
	db domain.IDonacion
}

func NewCreateDonacion(db domain.IDonacion) *CreateDonacion {
	return &CreateDonacion{db: db}
}

func (cd *CreateDonacion) Execute(donacion entities.Donacion) (entities.Donacion, error) {
	if donacion.UsuarioID <= 0 {
		return entities.Donacion{}, errors.New("el usuario_id es obligatorio")
	}
	if donacion.ModuloID <= 0 {
		return entities.Donacion{}, errors.New("el modulo_id es obligatorio")
	}
	if donacion.Monto <= 0 {
		return entities.Donacion{}, errors.New("el monto debe ser mayor a 0")
	}
	if donacion.Estado == "" {
		return entities.Donacion{}, errors.New("el estado es obligatorio")
	}

	// Establecer valores por defecto
	if donacion.Moneda == "" {
		donacion.Moneda = "MXN"
	}
	if donacion.FechaPago.IsZero() {
		donacion.FechaPago = time.Now()
	}

	return cd.db.Save(donacion)
}
