package infrastructure

import (
	"AmethToledo/src/core"
	"AmethToledo/src/donaciones/application"
	"AmethToledo/src/donaciones/infrastructure/adapters"
	"AmethToledo/src/donaciones/infrastructure/controllers"
)

type DependenciesDonaciones struct {
	CreateDonacionController         *controllers.CreateDonacionController
	GetAllDonacionesController       *controllers.GetAllDonacionesController
	GetDonacionByIdController        *controllers.GetDonacionByIdController
	GetDonacionesByUsuarioController *controllers.GetDonacionesByUsuarioController
	GetDonacionesByModuloController  *controllers.GetDonacionesByModuloController
	GetTotalDonacionesController     *controllers.GetTotalDonacionesController
	UpdateDonacionController         *controllers.UpdateDonacionController
	DeleteDonacionController         *controllers.DeleteDonacionController
	GetStatsDonacionesController     *controllers.GetStatsDonacionesController
}

func InitDonaciones() *DependenciesDonaciones {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	// Application layer
	createDonacionApp := application.NewCreateDonacion(ps)
	getAllDonacionesApp := application.NewGetAllDonaciones(ps)
	getDonacionByIdApp := application.NewGetDonacionById(ps)
	getDonacionesByUsuarioApp := application.NewGetDonacionesByUsuario(ps)
	getDonacionesByModuloApp := application.NewGetDonacionesByModulo(ps)
	getTotalDonacionesApp := application.NewGetTotalDonaciones(ps)
	updateDonacionApp := application.NewUpdateDonacion(ps)
	deleteDonacionApp := application.NewDeleteDonacion(ps)
	getStatsDonacionesApp := application.NewGetStatsDonaciones(ps)

	return &DependenciesDonaciones{
		CreateDonacionController:         controllers.NewCreateDonacionController(createDonacionApp),
		GetAllDonacionesController:       controllers.NewGetAllDonacionesController(getAllDonacionesApp),
		GetDonacionByIdController:        controllers.NewGetDonacionByIdController(getDonacionByIdApp),
		GetDonacionesByUsuarioController: controllers.NewGetDonacionesByUsuarioController(getDonacionesByUsuarioApp),
		GetDonacionesByModuloController:  controllers.NewGetDonacionesByModuloController(getDonacionesByModuloApp),
		GetTotalDonacionesController:     controllers.NewGetTotalDonacionesController(getTotalDonacionesApp),
		UpdateDonacionController:         controllers.NewUpdateDonacionController(updateDonacionApp),
		DeleteDonacionController:         controllers.NewDeleteDonacionController(deleteDonacionApp),
		GetStatsDonacionesController:     controllers.NewGetStatsDonacionesController(getStatsDonacionesApp),
	}
}
