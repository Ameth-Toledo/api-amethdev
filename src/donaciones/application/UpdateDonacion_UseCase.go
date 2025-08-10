package application

import (
	"AmethToledo/src/donaciones/domain"
	"AmethToledo/src/donaciones/domain/entities"
	"errors"
)

type UpdateDonacion struct {
	db domain.IDonacion
}

func NewUpdateDonacion(db domain.IDonacion) *UpdateDonacion {
	return &UpdateDonacion{db: db}
}

func (ud *UpdateDonacion) Execute(donacion entities.Donacion) error {
	if donacion.ID <= 0 {
		return errors.New("id inválido")
	}

	// Verificar que la donación existe
	existing, err := ud.db.GetById(donacion.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("donación no encontrada")
	}

	if donacion.UsuarioID <= 0 {
		return errors.New("el usuario_id es obligatorio")
	}
	if donacion.ModuloID <= 0 {
		return errors.New("el modulo_id es obligatorio")
	}
	if donacion.Monto <= 0 {
		return errors.New("el monto debe ser mayor a 0")
	}
	if donacion.Estado == "" {
		return errors.New("el estado es obligatorio")
	}

	// Establecer valores por defecto
	if donacion.Moneda == "" {
		donacion.Moneda = "MXN"
	}

	return ud.db.Update(donacion)
}
