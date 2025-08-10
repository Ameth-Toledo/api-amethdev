package routes

import (
	"AmethToledo/src/modulos/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureModuloRoutes(router *gin.Engine,
	createModuloController *controllers.CreateModuloController,
	getAllModulosController *controllers.GetAllModulosController,
	getByIdModuloController *controllers.GetModuloByIdController,
	updateModuloController *controllers.UpdateModuloController,
	deleteModuloController *controllers.DeleteModuloController,
	searchModulosController *controllers.SearchModulosController,
) {
	moduloGroup := router.Group("/modulos")
	{
		moduloGroup.POST("", createModuloController.Execute)        // POST /modulos
		moduloGroup.GET("", getAllModulosController.Execute)        // GET /modulos (con soporte para ?id= y ?titulo=)
		moduloGroup.GET("/search", searchModulosController.Execute) // GET /modulos/search?titulo=
		moduloGroup.GET("/:id", getByIdModuloController.Execute)
		moduloGroup.PUT("/:id", updateModuloController.Execute)
		moduloGroup.DELETE("/:id", deleteModuloController.Execute)
	}
}
