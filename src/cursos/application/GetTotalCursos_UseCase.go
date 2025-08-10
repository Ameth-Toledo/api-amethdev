package application

import "AmethToledo/src/cursos/domain"

type GetTotalCursos struct {
	db domain.ICurso
}

func NewGetTotalCursos(db domain.ICurso) *GetTotalCursos {
	return &GetTotalCursos{db: db}
}

func (gtc *GetTotalCursos) Execute() (int, error) {
	return gtc.db.GetTotal()
}
