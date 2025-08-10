package main

import (
	comentarioInfra "AmethToledo/src/comentarios/infrastructure"
	comentarioRoutes "AmethToledo/src/comentarios/infrastructure/routes"
	contenidoInfra "AmethToledo/src/contenidos/infrastructure"
	contenidoRoutes "AmethToledo/src/contenidos/infrastructure/routes"
	cursoInfra "AmethToledo/src/cursos/infrastructure"
	cursoRoutes "AmethToledo/src/cursos/infrastructure/routes"
	donacionInfra "AmethToledo/src/donaciones/infrastructure"
	donacionRoutes "AmethToledo/src/donaciones/infrastructure/routes"
	likeInfra "AmethToledo/src/likes/infrastructure"
	likeRoutes "AmethToledo/src/likes/infrastructure/routes"
	moduloInfra "AmethToledo/src/modulos/infrastructure"
	moduloRoutes "AmethToledo/src/modulos/infrastructure/routes"
	userInfra "AmethToledo/src/users/insfrastructure"
	userRoutes "AmethToledo/src/users/insfrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Inicializar dependencias de usuarios
	userDependencies := userInfra.InitUsers()
	userRoutes.ConfigureUserRoutes(r,
		userDependencies.CreateUserController,
		userDependencies.GetAllUsersController,
		userDependencies.GetByIdUserController,
		userDependencies.GetTotalUsersController,
		userDependencies.UpdateUserController,
		userDependencies.DeleteUserController,
		userDependencies.AuthController,
	)

	// Inicializar dependencias de cursos
	cursoDependencies := cursoInfra.InitCursos()
	cursoRoutes.ConfigureCursoRoutes(r,
		cursoDependencies.CreateCursoController,
		cursoDependencies.GetAllCursosController,
		cursoDependencies.GetByIdCursoController,
		cursoDependencies.UpdateCursoController,
		cursoDependencies.DeleteCursoController,
		cursoDependencies.SearchCursosController,
		cursoDependencies.GetTotalCursosController,
	)

	moduloDependencies := moduloInfra.InitModulos()
	moduloRoutes.ConfigureModuloRoutes(r,
		moduloDependencies.CreateModuloController,
		moduloDependencies.GetAllModuloController,
		moduloDependencies.GetByIdModuloController,
		moduloDependencies.UpdateModuloController,
		moduloDependencies.DeleteModuloController,
		moduloDependencies.SearchModuloController,
	)

	contenidoDependencies := contenidoInfra.InitContenidos()
	contenidoRoutes.ConfigureContenidoRoutes(r,
		contenidoDependencies.CreateContenidoController,
		contenidoDependencies.GetAllContenidosController,
		contenidoDependencies.GetByIdContenidoController,
		contenidoDependencies.UpdateContenidoController,
		contenidoDependencies.DeleteContenidoController,
	)

	likeDependencies := likeInfra.InitLikes()
	likeRoutes.ConfigureLikeRoutes(r,
		likeDependencies.ToggleLikeController,
		likeDependencies.GetLikeCountController,
		likeDependencies.GetLikesByModuloController,
		likeDependencies.GetMostLikedModulosController,
		likeDependencies.GetLikesByUserController,
		likeDependencies.GetLikeStatsController,
	)

	donacionDependencies := donacionInfra.InitDonaciones()
	donacionRoutes.ConfigureDonacionRoutes(r,
		donacionDependencies.CreateDonacionController,
		donacionDependencies.GetAllDonacionesController,
		donacionDependencies.GetDonacionByIdController,
		donacionDependencies.GetDonacionesByUsuarioController,
		donacionDependencies.GetDonacionesByModuloController,
		donacionDependencies.GetTotalDonacionesController,
		donacionDependencies.UpdateDonacionController,
		donacionDependencies.DeleteDonacionController,
		donacionDependencies.GetStatsDonacionesController,
	)

	comentarioDependencies := comentarioInfra.InitComentarios()
	comentarioRoutes.ConfigureComentarioRoutes(r,
		comentarioDependencies.CreateComentarioController,
		comentarioDependencies.GetAllComentariosController,
		comentarioDependencies.GetByIdComentarioController,
		comentarioDependencies.GetComentariosByModuloController,
		comentarioDependencies.GetComentariosByUsuarioController,
		comentarioDependencies.GetTotalComentariosController,
		comentarioDependencies.UpdateComentarioController,
		comentarioDependencies.DeleteComentarioController,
		comentarioDependencies.GetComentariosByModuloWithUserController,
		comentarioDependencies.GetAllComentariosWithUserController,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
