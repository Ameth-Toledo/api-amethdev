package application

import (
	"AmethToledo/src/cursos/domain"
	"AmethToledo/src/cursos/domain/entities"
	"errors"
)

type GetCursoById struct {
	db domain.ICurso
}

func NewGetCursoById(db domain.ICurso) *GetCursoById {
	return &GetCursoById{db: db}
}

func (gcb *GetCursoById) Execute(id int) (*entities.Curso, error) {
	if id <= 0 {
		return nil, errors.New("id inválido")
	}

	return gcb.db.GetById(id)
}
