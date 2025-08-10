package infrastructure

import (
	"AmethToledo/src/comentarios/application"
	"AmethToledo/src/comentarios/infrastructure/adapters"
	"AmethToledo/src/comentarios/infrastructure/controllers"
	"AmethToledo/src/core"
)

type DependenciesComentarios struct {
	CreateComentarioController        *controllers.CreateComentarioController
	GetAllComentariosController       *controllers.GetAllComentariosController
	GetByIdComentarioController       *controllers.GetComentarioByIdController
	GetComentariosByModuloController  *controllers.GetComentariosByModuloController
	GetComentariosByUsuarioController *controllers.GetComentariosByUsuarioController
	GetTotalComentariosController     *controllers.GetTotalComentariosController
	UpdateComentarioController        *controllers.UpdateComentarioController
	DeleteComentarioController        *controllers.DeleteComentarioController
	// Nuevos controladores
	GetComentariosByModuloWithUserController *controllers.GetComentariosByModuloWithUserController
	GetAllComentariosWithUserController      *controllers.GetAllComentariosWithUserController
}

func InitComentarios() *DependenciesComentarios {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	createComentarioApp := application.NewCreateComentario(ps)
	getAllComentariosApp := application.NewGetAllComentarios(ps)
	getComentarioByIdApp := application.NewGetComentarioById(ps)
	getComentariosByModuloApp := application.NewGetComentariosByModulo(ps)
	getComentariosByUsuarioApp := application.NewGetComentariosByUsuario(ps)
	getTotalComentariosApp := application.NewGetTotalComentarios(ps)
	updateComentarioApp := application.NewUpdateComentario(ps)
	deleteComentarioApp := application.NewDeleteComentario(ps)

	// Nuevas aplicaciones
	getComentariosByModuloWithUserApp := application.NewGetComentariosByModuloWithUser(ps)
	getAllComentariosWithUserApp := application.NewGetAllComentariosWithUser(ps)

	return &DependenciesComentarios{
		CreateComentarioController:               controllers.NewCreateComentarioController(createComentarioApp),
		GetAllComentariosController:              controllers.NewGetAllComentariosController(getAllComentariosApp),
		GetByIdComentarioController:              controllers.NewGetComentarioByIdController(getComentarioByIdApp),
		GetComentariosByModuloController:         controllers.NewGetComentariosByModuloController(getComentariosByModuloApp),
		GetComentariosByUsuarioController:        controllers.NewGetComentariosByUsuarioController(getComentariosByUsuarioApp),
		GetTotalComentariosController:            controllers.NewGetTotalComentariosController(getTotalComentariosApp),
		UpdateComentarioController:               controllers.NewUpdateComentarioController(updateComentarioApp),
		DeleteComentarioController:               controllers.NewDeleteComentarioController(deleteComentarioApp),
		GetComentariosByModuloWithUserController: controllers.NewGetComentariosByModuloWithUserController(getComentariosByModuloWithUserApp),
		GetAllComentariosWithUserController:      controllers.NewGetAllComentariosWithUserController(getAllComentariosWithUserApp),
	}
}
