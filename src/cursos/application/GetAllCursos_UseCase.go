package application

import (
	"AmethToledo/src/cursos/domain"
	"AmethToledo/src/cursos/domain/entities"
)

type GetAllCursos struct {
	db domain.ICurso
}

func NewGetAllCursos(db domain.ICurso) *GetAllCursos {
	return &GetAllCursos{db: db}
}

func (gac *GetAllCursos) Execute() ([]entities.Curso, error) {
	return gac.db.GetAll()
}
