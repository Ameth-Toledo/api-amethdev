package application

import (
	"AmethToledo/src/modulos/domain"
	"AmethToledo/src/modulos/domain/entities"
	"errors"
)

type SearchModulos struct {
	db domain.IModulo
}

func NewSearchModulos(db domain.IModulo) *SearchModulos {
	return &SearchModulos{db: db}
}

func (sm *SearchModulos) Execute(titulo string) ([]entities.Modulo, error) {
	if titulo == "" {
		return nil, errors.New("el título de búsqueda es obligatorio")
	}

	return sm.db.GetByTitulo(titulo)
}
