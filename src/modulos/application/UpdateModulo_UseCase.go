package application

import (
	"AmethToledo/src/modulos/domain"
	"AmethToledo/src/modulos/domain/entities"
	"errors"
)

type UpdateModulo struct {
	db domain.IModulo
}

func NewUpdateModulo(db domain.IModulo) *UpdateModulo {
	return &UpdateModulo{db: db}
}

func (um *UpdateModulo) Execute(modulo entities.Modulo) error {
	if modulo.ID <= 0 {
		return errors.New("id inválido")
	}
	if modulo.IdCurso <= 0 {
		return errors.New("el id del curso es obligatorio y debe ser válido")
	}
	if modulo.ImagenPortada == "" {
		return errors.New("la imagen de portada es obligatoria")
	}
	if modulo.Titulo == "" {
		return errors.New("el título es obligatorio")
	}
	if modulo.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}

	return um.db.Update(modulo)
}
