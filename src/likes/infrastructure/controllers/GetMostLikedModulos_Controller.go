package controllers

import (
	"AmethToledo/src/likes/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetMostLikedModulosController struct {
	getMostLikedModulos *application.GetMostLikedModulos
}

func NewGetMostLikedModulosController(getMostLikedModulos *application.GetMostLikedModulos) *GetMostLikedModulosController {
	return &GetMostLikedModulosController{
		getMostLikedModulos: getMostLikedModulos,
	}
}

func (gmlmc *GetMostLikedModulosController) Execute(c *gin.Context) {
	// Obtener limit del query parameter (opcional, default 10)
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	modulos, err := gmlmc.getMostLikedModulos.Execute(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Módulos más likeados obtenidos exitosamente",
		"modulos":       modulos,
		"total_modulos": len(modulos),
		"limit_applied": limit,
	})
}
