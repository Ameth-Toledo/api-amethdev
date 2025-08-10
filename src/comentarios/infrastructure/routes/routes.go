package routes

import (
	"AmethToledo/src/comentarios/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureComentarioRoutes(router *gin.Engine,
	createComentarioController *controllers.CreateComentarioController,
	getAllComentariosController *controllers.GetAllComentariosController,
	getByIdComentarioController *controllers.GetComentarioByIdController,
	getComentariosByModuloController *controllers.GetComentariosByModuloController,
	getComentariosByUsuarioController *controllers.GetComentariosByUsuarioController,
	getTotalComentariosController *controllers.GetTotalComentariosController,
	updateComentarioController *controllers.UpdateComentarioController,
	deleteComentarioController *controllers.DeleteComentarioController,
	getComentariosByModuloWithUserController *controllers.GetComentariosByModuloWithUserController,
	getAllComentariosWithUserController *controllers.GetAllComentariosWithUserController,
) {
	comentarioGroup := router.Group("/comentarios")
	{
		comentarioGroup.POST("", createComentarioController.Execute)                                         // POST /comentarios
		comentarioGroup.GET("", getAllComentariosController.Execute)                                         // GET /comentarios
		comentarioGroup.GET("/:id", getByIdComentarioController.Execute)                                     // GET /comentarios/:id
		comentarioGroup.GET("/total", getTotalComentariosController.Execute)                                 // GET /comentarios/total
		comentarioGroup.PUT("/:id", updateComentarioController.Execute)                                      // PUT /comentarios/:id
		comentarioGroup.DELETE("/:id", deleteComentarioController.Execute)                                   // DELETE /comentarios/:id
		comentarioGroup.GET("/modulo/:moduloId", getComentariosByModuloController.Execute)                   // GET /comentarios/modulo/:moduloId
		comentarioGroup.GET("/usuario/:usuarioId", getComentariosByUsuarioController.Execute)                // GET /comentarios/usuario/:usuarioId
		comentarioGroup.GET("/modulo/:moduloId/with-user", getComentariosByModuloWithUserController.Execute) // GET /comentarios/modulo/:moduloId/with-user
		comentarioGroup.GET("/with-user", getAllComentariosWithUserController.Execute)                       // GET /comentarios/with-user
	}
}
