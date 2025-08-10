package controllers

import (
	"AmethToledo/src/likes/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetLikeCountController struct {
	getLikeCount *application.GetLikeCount
}

func NewGetLikeCountController(getLikeCount *application.GetLikeCount) *GetLikeCountController {
	return &GetLikeCountController{
		getLikeCount: getLikeCount,
	}
}

func (glcc *GetLikeCountController) Execute(c *gin.Context) {
	// Obtener modulo_id del path parameter
	moduloIDStr := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	// Obtener parámetros opcionales del query string
	var usuarioID *int
	var fingerprintHash *string

	if usuarioIDStr := c.Query("usuario_id"); usuarioIDStr != "" {
		if id, err := strconv.Atoi(usuarioIDStr); err == nil {
			usuarioID = &id
		}
	}

	if fingerprint := c.Query("fingerprint_hash"); fingerprint != "" {
		fingerprintHash = &fingerprint
	}

	response, err := glcc.getLikeCount.Execute(moduloID, usuarioID, fingerprintHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
