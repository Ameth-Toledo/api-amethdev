package infrastructure

import (
	"AmethToledo/src/contenidos/application"
	"AmethToledo/src/contenidos/infrastructure/adapters"
	"AmethToledo/src/contenidos/infrastructure/controllers"
	"AmethToledo/src/core"
)

type DependenciesContenidos struct {
	CreateContenidoController  *controllers.CreateContenidoController
	GetAllContenidosController *controllers.GetAllContenidosController
	GetByIdContenidoController *controllers.GetContenidoByIdController
	UpdateContenidoController  *controllers.UpdateContenidoController
	DeleteContenidoController  *controllers.DeleteContenidoController
}

func InitContenidos() *DependenciesContenidos {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	createContenidoApp := application.NewCreateContenido(ps)
	getAllContenidosApp := application.NewGetAllContenidos(ps)
	getContenidoByIdApp := application.NewGetContenidoById(ps)
	updateContenidoApp := application.NewUpdateContenido(ps)
	deleteContenidoApp := application.NewDeleteContenido(ps)

	return &DependenciesContenidos{
		CreateContenidoController:  controllers.NewCreateContenidoController(createContenidoApp),
		GetAllContenidosController: controllers.NewGetAllContenidosController(getAllContenidosApp, getContenidoByIdApp),
		GetByIdContenidoController: controllers.NewGetContenidoByIdController(getContenidoByIdApp),
		UpdateContenidoController:  controllers.NewUpdateContenidoController(updateContenidoApp),
		DeleteContenidoController:  controllers.NewDeleteContenidoController(deleteContenidoApp),
	}
}
