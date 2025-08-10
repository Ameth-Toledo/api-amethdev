package domain

import "AmethToledo/src/users/domain/entities"

type IUser interface {
	Save(user entities.User) (entities.User, error)
	GetByCorreo(email string) (*entities.User, error)
	GetAll() ([]entities.User, error)
	GetById(id int) (*entities.User, error)
	GetTotal() (int, error)
	Update(user entities.User) error
	Delete(id int) error
}
