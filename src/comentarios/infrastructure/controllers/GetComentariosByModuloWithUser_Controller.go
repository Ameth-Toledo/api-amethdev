package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetComentariosByModuloWithUserController struct {
	getComentariosByModuloWithUser *application.GetComentariosByModuloWithUser
}

func NewGetComentariosByModuloWithUserController(getComentariosByModuloWithUser *application.GetComentariosByModuloWithUser) *GetComentariosByModuloWithUserController {
	return &GetComentariosByModuloWithUserController{getComentariosByModuloWithUser: getComentariosByModuloWithUser}
}

func (gcbmwuc *GetComentariosByModuloWithUserController) Execute(c *gin.Context) {
	moduloIdStr := c.Param("moduloId")
	moduloId, err := strconv.Atoi(moduloIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	comentarios, err := gcbmwuc.getComentariosByModuloWithUser.Execute(moduloId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comentarios": comentarios})
}
