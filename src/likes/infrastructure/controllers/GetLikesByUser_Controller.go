package controllers

import (
	"AmethToledo/src/likes/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetLikesByUserController struct {
	getLikesByUser *application.GetLikesByUser
}

func NewGetLikesByUserController(getLikesByUser *application.GetLikesByUser) *GetLikesByUserController {
	return &GetLikesByUserController{
		getLikesByUser: getLikesByUser,
	}
}

func (glbuc *GetLikesByUserController) Execute(c *gin.Context) {
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

	response, err := glbuc.getLikesByUser.Execute(usuarioID, fingerprintHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Likes del usuario obtenidos exitosamente",
		"data":    response,
	})
}
