package insfrastructure

import (
	"AmethToledo/src/core"
	"AmethToledo/src/users/application"
	"AmethToledo/src/users/insfrastructure/adapters"
	"AmethToledo/src/users/insfrastructure/controllers"
)

type DependenciesUsers struct {
	CreateUserController  *controllers.CreateUserController
	GetAllUsersController *controllers.GetAllUsersController
	GetByIdUserController *controllers.GetUserByIdController
	UpdateUserController  *controllers.UpdateUserController
	DeleteUserController  *controllers.DeleteUserController
	AuthController        *controllers.AuthController
}

func InitUsers() *DependenciesUsers {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	createUserApp := application.NewCreateUser(ps)
	getAllUsersApp := application.NewGetAllUsers(ps)
	getUserByIdApp := application.NewGetUserById(ps)
	updateUserApp := application.NewUpdateUser(ps)
	deleteUserApp := application.NewDeleteUser(ps)
	authService := application.NewAuthService(ps)

	return &DependenciesUsers{
		CreateUserController:  controllers.NewCreateUserController(createUserApp, authService),
		GetAllUsersController: controllers.NewGetAllUsersController(getAllUsersApp),
		GetByIdUserController: controllers.NewGetUserByIdController(getUserByIdApp),
		UpdateUserController:  controllers.NewUpdateUserController(updateUserApp),
		DeleteUserController:  controllers.NewDeleteUserController(deleteUserApp),
		AuthController:        controllers.NewAuthController(authService),
	}
}
