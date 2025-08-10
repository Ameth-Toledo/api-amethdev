package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetComentariosByModuloController struct {
	getComentariosByModulo *application.GetComentariosByModulo
}

func NewGetComentariosByModuloController(getComentariosByModulo *application.GetComentariosByModulo) *GetComentariosByModuloController {
	return &GetComentariosByModuloController{getComentariosByModulo: getComentariosByModulo}
}

func (gcbmc *GetComentariosByModuloController) Execute(c *gin.Context) {
	moduloIdStr := c.Param("moduloId")
	moduloId, err := strconv.Atoi(moduloIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	comentarios, err := gcbmc.getComentariosByModulo.Execute(moduloId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comentarios": comentarios})
}
