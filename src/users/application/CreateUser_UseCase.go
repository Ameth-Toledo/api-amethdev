package application

import (
	"AmethToledo/src/users/domain"
	"AmethToledo/src/users/domain/entities"
	"errors"
	"time"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cu *CreateUser) Execute(user entities.User) (entities.User, error) {
	// Validaciones básicas
	if user.Nombres == "" {
		return entities.User{}, errors.New("el nombre es obligatorio")
	}
	if user.Email == "" {
		return entities.User{}, errors.New("el email es obligatorio")
	}
	if user.PasswordHash == "" {
		return entities.User{}, errors.New("la contraseña es obligatoria")
	}

	// Verificar si el email ya existe
	existingUser, _ := cu.db.GetByCorreo(user.Email)
	if existingUser != nil {
		return entities.User{}, errors.New("el email ya está registrado")
	}

	// Establecer fecha de registro
	user.FechaRegistro = time.Now()

	return cu.db.Save(user)
}
