package infrastructure

import (
	"AmethToledo/src/core"
	"AmethToledo/src/modulos/application"
	"AmethToledo/src/modulos/infrastructure/adapters"
	"AmethToledo/src/modulos/infrastructure/controllers"
)

type DependenciesModulos struct {
	CreateModuloController  *controllers.CreateModuloController
	GetAllModuloController  *controllers.GetAllModulosController
	GetByIdModuloController *controllers.GetModuloByIdController
	UpdateModuloController  *controllers.UpdateModuloController
	DeleteModuloController  *controllers.DeleteModuloController
	SearchModuloController  *controllers.SearchModulosController
}

func InitModulos() *DependenciesModulos {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	createModuloApp := application.NewCreateModulo(ps)
	getAllModuloApp := application.NewGetAllModulos(ps)
	getByIdModuloApp := application.NewGetModuloById(ps)
	updateModuloApp := application.NewUpdateModulo(ps)
	deleteModuloApp := application.NewDeleteModulo(ps)
	searchModuloApp := application.NewSearchModulos(ps)

	return &DependenciesModulos{
		CreateModuloController:  controllers.NewCreateModuloController(createModuloApp),
		GetAllModuloController:  controllers.NewGetAllModulosController(getAllModuloApp, getByIdModuloApp, searchModuloApp),
		GetByIdModuloController: controllers.NewGetModuloByIdController(getByIdModuloApp),
		UpdateModuloController:  controllers.NewUpdateModuloController(updateModuloApp),
		DeleteModuloController:  controllers.NewDeleteModuloController(deleteModuloApp),
		SearchModuloController:  controllers.NewSearchModulosController(searchModuloApp),
	}
}
