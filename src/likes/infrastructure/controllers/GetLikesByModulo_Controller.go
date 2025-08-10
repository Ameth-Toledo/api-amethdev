package controllers

import (
	"AmethToledo/src/likes/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetLikesByModuloController struct {
	getLikesByModulo *application.GetLikesByModulo
}

func NewGetLikesByModuloController(getLikesByModulo *application.GetLikesByModulo) *GetLikesByModuloController {
	return &GetLikesByModuloController{
		getLikesByModulo: getLikesByModulo,
	}
}

func (glbmc *GetLikesByModuloController) Execute(c *gin.Context) {
	// Obtener modulo_id del path parameter
	moduloIDStr := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	likes, err := glbmc.getLikesByModulo.Execute(moduloID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"modulo_id": moduloID,
		"likes":     likes,
		"total":     len(likes),
	})
}
