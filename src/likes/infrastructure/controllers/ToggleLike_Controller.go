package controllers

import (
	"AmethToledo/src/likes/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ToggleLikeController struct {
	toggleLike *application.ToggleLike
}

func NewToggleLikeController(toggleLike *application.ToggleLike) *ToggleLikeController {
	return &ToggleLikeController{
		toggleLike: toggleLike,
	}
}

type ToggleLikeRequest struct {
	UsuarioID       *int    `json:"usuario_id,omitempty"`
	FingerprintHash *string `json:"fingerprint_hash,omitempty"`
}

func (tlc *ToggleLikeController) Execute(c *gin.Context) {
	// Obtener modulo_id del path parameter
	moduloIDStr := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del módulo inválido"})
		return
	}

	var req ToggleLikeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toggleReq := application.ToggleLikeRequest{
		ModuloID:        moduloID,
		UsuarioID:       req.UsuarioID,
		FingerprintHash: req.FingerprintHash,
	}

	response, err := tlc.toggleLike.Execute(toggleReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
