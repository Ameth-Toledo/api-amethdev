package main

import (
	userInfra "AmethToledo/src/users/insfrastructure"
	userRoutes "AmethToledo/src/users/insfrastructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	//	r.StaticFS("/img", gin.Dir("./static/img", false))
	//  r.StaticFS("/uploads", gin.Dir("./uploads", false))

	userDependencies := userInfra.InitUsers()
	userRoutes.ConfigureUserRoutes(r,
		userDependencies.CreateUserController,
		userDependencies.GetAllUsersController,
		userDependencies.GetByIdUserController,
		userDependencies.UpdateUserController,
		userDependencies.DeleteUserController,
		userDependencies.AuthController,
	)

	r.Run(":8080")
}
