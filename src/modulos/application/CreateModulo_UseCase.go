package application

import (
	"AmethToledo/src/modulos/domain"
	"AmethToledo/src/modulos/domain/entities"
	"errors"
)

type CreateModulo struct {
	db domain.IModulo
}

func NewCreateModulo(db domain.IModulo) *CreateModulo {
	return &CreateModulo{db: db}
}

func (cm *CreateModulo) Execute(modulo entities.Modulo) (entities.Modulo, error) {
	if modulo.IdCurso <= 0 {
		return entities.Modulo{}, errors.New("el id del curso es obligatorio y debe ser válido")
	}
	if modulo.ImagenPortada == "" {
		return entities.Modulo{}, errors.New("la imagen de portada es obligatoria")
	}
	if modulo.Titulo == "" {
		return entities.Modulo{}, errors.New("el título es obligatorio")
	}
	if modulo.Descripcion == "" {
		return entities.Modulo{}, errors.New("la descripción es obligatoria")
	}

	return cm.db.Save(modulo)
}
