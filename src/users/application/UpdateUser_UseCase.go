package application

import (
	"AmethToledo/src/users/domain"
	"AmethToledo/src/users/domain/entities"
	"errors"
)

type UpdateUser struct {
	db domain.IUser
}

func NewUpdateUser(db domain.IUser) *UpdateUser {
	return &UpdateUser{db: db}
}

func (uu *UpdateUser) Execute(user entities.User) error {
	if user.ID <= 0 {
		return errors.New("ID inválido")
	}

	// Verificar que el usuario existe
	existingUser, err := uu.db.GetById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("usuario no encontrado")
	}

	return uu.db.Update(user)
}
