package routes

import (
	"AmethToledo/src/cursos/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureCursoRoutes(router *gin.Engine,
	createCursoController *controllers.CreateCursoController,
	getAllCursosController *controllers.GetAllCursosController,
	getByIdCursoController *controllers.GetCursoByIdController,
	updateCursoController *controllers.UpdateCursoController,
	deleteCursoController *controllers.DeleteCursoController,
	searchCursosController *controllers.SearchCursosController,
	getTotalCursosController *controllers.GetTotalCursosController,
) {
	// Rutas para cursos
	cursoGroup := router.Group("/cursos")
	{
		cursoGroup.POST("", createCursoController.Execute)         // POST /cursos
		cursoGroup.GET("", getAllCursosController.Execute)         // GET /cursos (con soporte para ?id= y ?nombre=)
		cursoGroup.GET("/search", searchCursosController.Execute)  // GET /cursos/search?nombre=
		cursoGroup.GET("/total", getTotalCursosController.Execute) // GET /cursos/total
		cursoGroup.GET("/:id", getByIdCursoController.Execute)     // GET /cursos/:id
		cursoGroup.PUT("/:id", updateCursoController.Execute)      // PUT /cursos/:id
		cursoGroup.DELETE("/:id", deleteCursoController.Execute)   // DELETE /cursos/:id
	}
}
