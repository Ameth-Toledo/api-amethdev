package application

import (
	"AmethToledo/src/users/domain"
	"AmethToledo/src/users/domain/entities"
	"errors"
)

type GetUserById struct {
	db domain.IUser
}

func NewGetUserById(db domain.IUser) *GetUserById {
	return &GetUserById{db: db}
}

func (gu *GetUserById) Execute(id int) (*entities.User, error) {
	if id <= 0 {
		return nil, errors.New("ID inválido")
	}
	return gu.db.GetById(id)
}
