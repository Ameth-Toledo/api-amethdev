package domain

import "AmethToledo/src/donaciones/domain/entities"

type IDonacion interface {
	Save(donacion entities.Donacion) (entities.Donacion, error)
	GetById(id int) (*entities.Donacion, error)
	GetAll() ([]entities.Donacion, error)
	GetByUsuarioID(usuarioID int) ([]entities.Donacion, error)
	GetByModuloID(moduloID int) ([]entities.Donacion, error)
	GetByEstado(estado string) ([]entities.Donacion, error)
	GetTotal() (int, error)
	GetTotalByUsuario(usuarioID int) (float64, error)
	GetTotalByModulo(moduloID int) (float64, error)
	Update(donacion entities.Donacion) error
	Delete(id int) error
}
