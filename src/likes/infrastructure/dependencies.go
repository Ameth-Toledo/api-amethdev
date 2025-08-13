package infrastructure

import (
	"AmethToledo/src/core"
	"AmethToledo/src/likes/application"
	"AmethToledo/src/likes/infrastructure/adapters"
	"AmethToledo/src/likes/infrastructure/controllers"
)

type DependenciesLikes struct {
	ToggleLikeController          *controllers.ToggleLikeController
	GetLikeCountController        *controllers.GetLikeCountController
	GetLikesByModuloController    *controllers.GetLikesByModuloController
	GetMostLikedModulosController *controllers.GetMostLikedModulosController
	GetLikesByUserController      *controllers.GetLikesByUserController
	GetLikeStatsController        *controllers.GetLikeStatsController

	GetAllModulosController                *controllers.GetAllModulosController
	GetLikesByModuloWithUserInfoController *controllers.GetLikesByModuloWithUserInfoController
	GetLikesByUserWithModuleInfoController *controllers.GetLikesByUserWithModuleInfoController
}

func InitLikes() *DependenciesLikes {
	conn := core.GetDBPool()
	ps := adapters.NewPostgreSQL(conn.DB)

	toggleLikeApp := application.NewToggleLike(ps)
	getLikeCountApp := application.NewGetLikeCount(ps)
	getLikesByModuloApp := application.NewGetLikesByModulo(ps)
	getMostLikedModulosApp := application.NewGetMostLikedModulos(ps)
	getLikesByUserApp := application.NewGetLikesByUser(ps)
	getLikeStatsApp := application.NewGetLikeStats(ps)

	getAllModulosApp := application.NewGetAllModulosUseCase(ps)
	getLikesByModuloWithUserInfoApp := application.NewGetLikesByModuloWithUserInfoUseCase(ps)
	getLikesByUserWithModuleInfoApp := application.NewGetLikesByUserWithModuleInfoUseCase(ps)

	return &DependenciesLikes{
		ToggleLikeController:                   controllers.NewToggleLikeController(toggleLikeApp),
		GetLikeCountController:                 controllers.NewGetLikeCountController(getLikeCountApp),
		GetLikesByModuloController:             controllers.NewGetLikesByModuloController(getLikesByModuloApp),
		GetMostLikedModulosController:          controllers.NewGetMostLikedModulosController(getMostLikedModulosApp),
		GetLikesByUserController:               controllers.NewGetLikesByUserController(getLikesByUserApp),
		GetLikeStatsController:                 controllers.NewGetLikeStatsController(getLikeStatsApp),
		GetAllModulosController:                controllers.NewGetAllModulosController(getAllModulosApp),
		GetLikesByModuloWithUserInfoController: controllers.NewGetLikesByModuloWithUserInfoController(getLikesByModuloWithUserInfoApp),
		GetLikesByUserWithModuleInfoController: controllers.NewGetLikesByUserWithModuleInfoController(getLikesByUserWithModuleInfoApp),
	}
}
