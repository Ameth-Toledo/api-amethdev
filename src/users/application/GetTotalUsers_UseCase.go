package application

import "AmethToledo/src/users/domain"

type GetTotalUsers struct {
	db domain.IUser
}

func NewGetTotalUsers(db domain.IUser) *GetTotalUsers {
	return &GetTotalUsers{db: db}
}

func (gtu *GetTotalUsers) Execute() (int, error) {
	return gtu.db.GetTotal()
}
