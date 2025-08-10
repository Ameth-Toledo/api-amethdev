package application

import (
	"AmethToledo/src/cursos/domain"
	"AmethToledo/src/cursos/domain/entities"
	"errors"
)

type SearchCursos struct {
	db domain.ICurso
}

func NewSearchCursos(db domain.ICurso) *SearchCursos {
	return &SearchCursos{db: db}
}

func (sc *SearchCursos) Execute(nombre string) ([]entities.Curso, error) {
	if nombre == "" {
		return nil, errors.New("el nombre de búsqueda es obligatorio")
	}

	return sc.db.GetByNombre(nombre)
}
