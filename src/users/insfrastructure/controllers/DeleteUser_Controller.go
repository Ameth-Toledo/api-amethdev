package controllers

import (
	"AmethToledo/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteUserController struct {
	deleteUser *application.DeleteUser
}

func NewDeleteUserController(deleteUser *application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{deleteUser: deleteUser}
}

func (dc *DeleteUserController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dc.deleteUser.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}
