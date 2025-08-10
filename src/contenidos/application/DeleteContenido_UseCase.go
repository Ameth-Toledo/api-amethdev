package application

import (
	"AmethToledo/src/contenidos/domain"
	"errors"
)

type DeleteContenido struct {
	db domain.IContenido
}

func NewDeleteContenido(db domain.IContenido) *DeleteContenido {
	return &DeleteContenido{db: db}
}

func (dc *DeleteContenido) Execute(id int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	return dc.db.Delete(id)
}
