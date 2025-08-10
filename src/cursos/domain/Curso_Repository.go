package domain

import "AmethToledo/src/cursos/domain/entities"

type ICurso interface {
	Save(curso entities.Curso) (entities.Curso, error)
	GetAll() ([]entities.Curso, error)
	GetById(id int) (*entities.Curso, error)
	GetByNombre(nombre string) ([]entities.Curso, error)
	GetTotal() (int, error)
	Update(curso entities.Curso) error
	Delete(id int) error
}
