package domain

import "AmethToledo/src/contenidos/domain/entities"

type IContenido interface {
	Save(contenido entities.Contenido) (entities.Contenido, error)
	GetAll() ([]entities.Contenido, error)
	GetById(id int) (*entities.Contenido, error)
	Update(contenido entities.Contenido) error
	Delete(id int) error
}
