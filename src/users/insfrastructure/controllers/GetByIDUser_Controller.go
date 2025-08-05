package controllers

import (
	"AmethToledo/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetUserByIdController struct {
	getUserById *application.GetUserById
}

func NewGetUserByIdController(getUserById *application.GetUserById) *GetUserByIdController {
	return &GetUserByIdController{getUserById: getUserById}
}

func (gc *GetUserByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	user, err := gc.getUserById.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Respuesta sin mostrar password_hash
	response := gin.H{
		"id":               user.ID,
		"nombres":          user.Nombres,
		"apellido_paterno": user.ApellidoPaterno,
		"apellido_materno": user.ApellidoMaterno,
		"email":            user.Email,
		"rol_id":           user.RolID,
		"avatar":           user.Avatar,
		"fecha_registro":   user.FechaRegistro,
	}

	c.JSON(http.StatusOK, gin.H{"user": response})
}
