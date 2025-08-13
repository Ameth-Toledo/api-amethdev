package routes

import (
	"AmethToledo/src/likes/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureLikeRoutes(router *gin.Engine,
	// Controladores existentes
	toggleLikeController *controllers.ToggleLikeController,
	getLikeCountController *controllers.GetLikeCountController,
	getLikesByModuloController *controllers.GetLikesByModuloController,
	getMostLikedModulosController *controllers.GetMostLikedModulosController,
	getLikesByUserController *controllers.GetLikesByUserController,
	getLikeStatsController *controllers.GetLikeStatsController,
	// Controladores nuevos
	getAllModulosController *controllers.GetAllModulosController,
	getLikesByModuloWithUserInfoController *controllers.GetLikesByModuloWithUserInfoController,
	getLikesByUserWithModuleInfoController *controllers.GetLikesByUserWithModuleInfoController,
) {
	likeGroup := router.Group("/likes")
	{
		// Rutas existentes
		likeGroup.POST("/modulo/:modulo_id/toggle", toggleLikeController.Execute)   // POST /likes/modulo/1/toggle
		likeGroup.GET("/modulo/:modulo_id", getLikeCountController.Execute)         // GET /likes/modulo/1?usuario_id=1&fingerprint_hash=abc123
		likeGroup.GET("/modulo/:modulo_id/all", getLikesByModuloController.Execute) // GET /likes/modulo/1/all
		likeGroup.GET("/modulos/most-liked", getMostLikedModulosController.Execute) // GET /likes/modulos/most-liked?limit=10
		likeGroup.GET("/user", getLikesByUserController.Execute)                    // GET /likes/user?usuario_id=1 o ?fingerprint_hash=abc123
		likeGroup.GET("/modulo/:modulo_id/stats", getLikeStatsController.Execute)   // GET /likes/modulo/1/stats?start_date=2024-01-01&end_date=2024-01-31

		// Nuevas rutas
		likeGroup.GET("/modulos/options", getAllModulosController.Execute)                              // GET /likes/modulos/options - Para selectores
		likeGroup.GET("/modulo/:modulo_id/detailed", getLikesByModuloWithUserInfoController.Execute)   // GET /likes/modulo/1/detailed - Likes con info de usuario
		likeGroup.GET("/user/detailed", getLikesByUserWithModuleInfoController.Execute)                // GET /likes/user/detailed?usuario_id=1 - Likes del usuario con info de módulo
	}
}