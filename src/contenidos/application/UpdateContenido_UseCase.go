package application

import (
	"AmethToledo/src/contenidos/domain"
	"AmethToledo/src/contenidos/domain/entities"
	"errors"
)

type UpdateContenido struct {
	db domain.IContenido
}

func NewUpdateContenido(db domain.IContenido) *UpdateContenido {
	return &UpdateContenido{db: db}
}

func (uc *UpdateContenido) Execute(contenido entities.Contenido) error {
	if contenido.ID <= 0 {
		return errors.New("id inválido")
	}
	if contenido.IdModulo <= 0 {
		return errors.New("el id del módulo es obligatorio y debe ser válido")
	}
	if contenido.ImagenPortada == "" {
		return errors.New("la imagen de portada es obligatoria")
	}
	if contenido.Titulo == "" {
		return errors.New("el título es obligatorio")
	}
	if contenido.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}

	if contenido.VideoURL == "" {
		return errors.New("la video url obligatoria")
	}

	if contenido.DescripcionModule == "" {
		return errors.New("la descripción del módulo es obligatoria")
	}

	return uc.db.Update(contenido)
}
