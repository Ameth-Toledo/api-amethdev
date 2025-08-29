package infrastructure

import (
	"AmethToledo/src/core"
	"AmethToledo/src/cursos/application"
	"AmethToledo/src/cursos/infrastructure/adapters"
	"AmethToledo/src/cursos/infrastructure/controllers"
	"AmethToledo/src/notifications"
)

type DependenciesCursos struct {
	CreateCursoController    *controllers.CreateCursoController
	GetAllCursosController   *controllers.GetAllCursosController
	GetByIdCursoController   *controllers.GetCursoByIdController
	UpdateCursoController    *controllers.UpdateCursoController
	DeleteCursoController    *controllers.DeleteCursoController
	SearchCursosController   *controllers.SearchCursosController
	GetTotalCursosController *controllers.GetTotalCursosController
	Hub                      *notifications.Hub // Agregar el hub
}

func InitCursos(hub *notifications.Hub) *DependenciesCursos {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	createCursoApp := application.NewCreateCurso(ps)
	getAllCursosApp := application.NewGetAllCursos(ps)
	getCursoByIdApp := application.NewGetCursoById(ps)
	updateCursoApp := application.NewUpdateCurso(ps)
	deleteCursoApp := application.NewDeleteCurso(ps)
	searchCursosApp := application.NewSearchCursos(ps)
	getTotalCursosApp := application.NewGetTotalCursos(ps)

	return &DependenciesCursos{
		CreateCursoController:    controllers.NewCreateCursoController(createCursoApp, hub),
		GetAllCursosController:   controllers.NewGetAllCursosController(getAllCursosApp, getCursoByIdApp, searchCursosApp),
		GetByIdCursoController:   controllers.NewGetCursoByIdController(getCursoByIdApp),
		UpdateCursoController:    controllers.NewUpdateCursoController(updateCursoApp),
		DeleteCursoController:    controllers.NewDeleteCursoController(deleteCursoApp),
		SearchCursosController:   controllers.NewSearchCursosController(searchCursosApp),
		GetTotalCursosController: controllers.NewGetTotalCursosController(getTotalCursosApp),
		Hub:                      hub,
	}
}
