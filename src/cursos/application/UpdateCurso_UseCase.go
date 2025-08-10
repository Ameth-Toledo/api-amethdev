package application

import (
	"AmethToledo/src/cursos/domain"
	"AmethToledo/src/cursos/domain/entities"
	"errors"
)

type UpdateCurso struct {
	db domain.ICurso
}

func NewUpdateCurso(db domain.ICurso) *UpdateCurso {
	return &UpdateCurso{db: db}
}

func (uc *UpdateCurso) Execute(curso entities.Curso) error {
	if curso.ID <= 0 {
		return errors.New("id inválido")
	}
	if curso.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}
	if curso.Nivel == "" {
		return errors.New("el nivel es obligatorio")
	}
	if curso.Duracion == "" {
		return errors.New("la duración es obligatoria")
	}
	if curso.Tecnologia == "" {
		return errors.New("la tecnología es obligatoria")
	}
	if curso.Fecha == "" {
		return errors.New("la fecha es obligatoria")
	}
	if curso.Imagen == "" {
		return errors.New("la imagen es obligatoria")
	}
	if curso.Descripcion == "" {
		return errors.New("la descripción es obligatoria")
	}

	return uc.db.Update(curso)
}
