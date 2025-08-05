package application

import (
	"AmethToledo/src/users/domain"
	"AmethToledo/src/users/domain/entities"
)

type GetAllUsers struct {
	db domain.IUser
}

func NewGetAllUsers(db domain.IUser) *GetAllUsers {
	return &GetAllUsers{db: db}
}

func (gu *GetAllUsers) Execute() ([]entities.User, error) {
	return gu.db.GetAll()
}
