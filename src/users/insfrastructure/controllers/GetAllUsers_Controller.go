package controllers

import (
	"AmethToledo/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllUsersController struct {
	getAllUsers *application.GetAllUsers
}

func NewGetAllUsersController(getAllUsers *application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{getAllUsers: getAllUsers}
}

func (gc *GetAllUsersController) Execute(c *gin.Context) {
	users, err := gc.getAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []gin.H
	for _, user := range users {
		response = append(response, gin.H{
			"id":               user.ID,
			"nombres":          user.Nombres,
			"apellido_paterno": user.ApellidoPaterno,
			"apellido_materno": user.ApellidoMaterno,
			"email":            user.Email,
			"rol_id":           user.RolID,
			"avatar":           user.Avatar,
			"fecha_registro":   user.FechaRegistro,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"users": response,
		"total": len(users),
	})
}
