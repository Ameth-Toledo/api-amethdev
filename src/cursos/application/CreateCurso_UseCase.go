package application

import (
	"AmethToledo/src/cursos/domain"
	"AmethToledo/src/cursos/domain/entities"
	"errors"
)

type CreateCurso struct {
	db domain.ICurso
}

func NewCreateCurso(db domain.ICurso) *CreateCurso {
	return &CreateCurso{db: db}
}

func (cc *CreateCurso) Execute(curso entities.Curso) (entities.Curso, error) {
	if curso.Nombre == "" {
		return entities.Curso{}, errors.New("el nombre es obligatorio")
	}
	if curso.Nivel == "" {
		return entities.Curso{}, errors.New("el nivel es obligatorio")
	}
	if curso.Duracion == "" {
		return entities.Curso{}, errors.New("la duración es obligatoria")
	}
	if curso.Tecnologia == "" {
		return entities.Curso{}, errors.New("la tecnología es obligatoria")
	}
	if curso.Fecha == "" {
		return entities.Curso{}, errors.New("la fecha es obligatoria")
	}
	if curso.Imagen == "" {
		return entities.Curso{}, errors.New("la imagen es obligatoria")
	}
	if curso.Descripcion == "" {
		return entities.Curso{}, errors.New("la descripción es obligatoria")
	}

	return cc.db.Save(curso)
}
