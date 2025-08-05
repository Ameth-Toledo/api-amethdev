package application

import (
	"AmethToledo/src/users/domain"
	"errors"
)

type DeleteUser struct {
	db domain.IUser
}

func NewDeleteUser(db domain.IUser) *DeleteUser {
	return &DeleteUser{db: db}
}

func (du *DeleteUser) Execute(id int) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	// Verificar que el usuario existe
	existingUser, err := du.db.GetById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	return du.db.Delete(id)
}
