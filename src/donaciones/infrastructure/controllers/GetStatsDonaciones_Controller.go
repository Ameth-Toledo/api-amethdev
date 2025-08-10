package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetStatsDonacionesController struct {
	getStatsDonaciones *application.GetStatsDonaciones
}

func NewGetStatsDonacionesController(getStatsDonaciones *application.GetStatsDonaciones) *GetStatsDonacionesController {
	return &GetStatsDonacionesController{
		getStatsDonaciones: getStatsDonaciones,
	}
}

func (gsdc *GetStatsDonacionesController) ExecuteByUsuario(c *gin.Context) {
	usuarioIDParam := c.Param("usuario_id")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario_id inválido"})
		return
	}

	total, err := gsdc.getStatsDonaciones.ExecuteByUsuario(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"usuario_id":   usuarioID,
		"total_donado": total,
		"moneda":       "MXN",
	})
}

func (gsdc *GetStatsDonacionesController) ExecuteByModulo(c *gin.Context) {
	moduloIDParam := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "modulo_id inválido"})
		return
	}

	total, err := gsdc.getStatsDonaciones.ExecuteByModulo(moduloID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"modulo_id":       moduloID,
		"total_recaudado": total,
		"moneda":          "MXN",
	})
}
