package routes

import (
	"AmethToledo/src/donaciones/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureDonacionRoutes(router *gin.Engine,
	createDonacionController *controllers.CreateDonacionController,
	getAllDonacionesController *controllers.GetAllDonacionesController,
	getDonacionByIdController *controllers.GetDonacionByIdController,
	getDonacionesByUsuarioController *controllers.GetDonacionesByUsuarioController,
	getDonacionesByModuloController *controllers.GetDonacionesByModuloController,
	getTotalDonacionesController *controllers.GetTotalDonacionesController,
	updateDonacionController *controllers.UpdateDonacionController,
	deleteDonacionController *controllers.DeleteDonacionController,
	getStatsDonacionesController *controllers.GetStatsDonacionesController,
) {
	donacionGroup := router.Group("/donaciones")
	{
		donacionGroup.POST("", createDonacionController.Execute)                            // POST /donaciones
		donacionGroup.GET("", getAllDonacionesController.Execute)                           // GET /donaciones
		donacionGroup.GET("/:id", getDonacionByIdController.Execute)                        // GET /donaciones/:id
		donacionGroup.GET("/usuario/:usuario_id", getDonacionesByUsuarioController.Execute) // GET /donaciones/usuario/:usuario_id
		donacionGroup.GET("/modulo/:modulo_id", getDonacionesByModuloController.Execute)    // GET /donaciones/modulo/:modulo_id
		donacionGroup.GET("/total", getTotalDonacionesController.Execute)                   // GET /donaciones/total
		donacionGroup.PUT("/:id", updateDonacionController.Execute)                         // PUT /donaciones/:id
		donacionGroup.DELETE("/:id", deleteDonacionController.Execute)                      // DELETE /donaciones/:id
	}

	statsGroup := router.Group("/stats/donaciones")
	{
		statsGroup.GET("/usuario/:usuario_id", getStatsDonacionesController.ExecuteByUsuario) // GET /stats/donaciones/usuario/:usuario_id
		statsGroup.GET("/modulo/:modulo_id", getStatsDonacionesController.ExecuteByModulo)    // GET /stats/donaciones/modulo/:modulo_id
	}
}
