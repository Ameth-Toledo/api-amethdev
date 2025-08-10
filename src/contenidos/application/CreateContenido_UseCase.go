package application

import (
	"AmethToledo/src/contenidos/domain"
	"AmethToledo/src/contenidos/domain/entities"
	"errors"
)

type CreateContenido struct {
	db domain.IContenido
}

func NewCreateContenido(db domain.IContenido) *CreateContenido {
	return &CreateContenido{db: db}
}

func (cc *CreateContenido) Execute(contenido entities.Contenido) (entities.Contenido, error) {
	if contenido.IdModulo <= 0 {
		return entities.Contenido{}, errors.New("el id del módulo es obligatorio y debe ser válido")
	}
	if contenido.ImagenPortada == "" {
		return entities.Contenido{}, errors.New("la imagen de portada es obligatoria")
	}
	if contenido.Titulo == "" {
		return entities.Contenido{}, errors.New("el título es obligatorio")
	}
	if contenido.Descripcion == "" {
		return entities.Contenido{}, errors.New("la descripción es obligatoria")
	}

	if contenido.VideoURL == "" {
		return entities.Contenido{}, errors.New("la video url obligatoria")
	}

	if contenido.DescripcionModule == "" {
		return entities.Contenido{}, errors.New("la descripción del módulo es obligatoria")
	}

	return cc.db.Save(contenido)
}
