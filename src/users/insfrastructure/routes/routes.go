package routes

import (
	"AmethToledo/src/users/insfrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureUserRoutes(router *gin.Engine,
	createUserController *controllers.CreateUserController,
	getAllUsersController *controllers.GetAllUsersController,
	getByIdUserController *controllers.GetUserByIdController,
	updateUserController *controllers.UpdateUserController,
	deleteUserController *controllers.DeleteUserController,
	authUserController *controllers.AuthController,
) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", createUserController.Execute)       // POST /users
		userGroup.GET("", getAllUsersController.Execute)       // GET /users
		userGroup.GET("/:id", getByIdUserController.Execute)   // GET /users/:id
		userGroup.PUT("/:id", updateUserController.Execute)    // PUT /users/:id
		userGroup.DELETE("/:id", deleteUserController.Execute) // DELETE /users/:id
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authUserController.Execute) // POST /auth/login
	}
}
