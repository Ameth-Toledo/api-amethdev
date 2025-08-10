package application

import (
	"AmethToledo/src/cursos/domain"
	"errors"
)

type DeleteCurso struct {
	db domain.ICurso
}

func NewDeleteCurso(db domain.ICurso) *DeleteCurso {
	return &DeleteCurso{db: db}
}

func (dc *DeleteCurso) Execute(id int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	return dc.db.Delete(id)
}
