package routes

import (
	"AmethToledo/src/contenidos/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureContenidoRoutes(router *gin.Engine,
	createContenidoController *controllers.CreateContenidoController,
	getAllContenidosController *controllers.GetAllContenidosController,
	getByIdContenidoController *controllers.GetContenidoByIdController,
	updateContenidoController *controllers.UpdateContenidoController,
	deleteContenidoController *controllers.DeleteContenidoController,
) {
	contenidoGroup := router.Group("/contenido")
	{
		contenidoGroup.POST("", createContenidoController.Execute)       // POST /contenido
		contenidoGroup.GET("", getAllContenidosController.Execute)       // GET /contenido
		contenidoGroup.GET("/:id", getByIdContenidoController.Execute)   // GET /contenido/:id
		contenidoGroup.PUT("/:id", updateContenidoController.Execute)    // PUT /contenido/:id
		contenidoGroup.DELETE("/:id", deleteContenidoController.Execute) // DELETE /contenido/:id
	}
}
