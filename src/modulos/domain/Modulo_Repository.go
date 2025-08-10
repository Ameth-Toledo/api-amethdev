package domain

import "AmethToledo/src/modulos/domain/entities"

type IModulo interface {
	Save(modulo entities.Modulo) (entities.Modulo, error)
	GetAll() ([]entities.Modulo, error)
	GetById(id int) (*entities.Modulo, error)
	GetByTitulo(titulo string) ([]entities.Modulo, error)
	Update(modulo entities.Modulo) error
	Delete(id int) error
}
