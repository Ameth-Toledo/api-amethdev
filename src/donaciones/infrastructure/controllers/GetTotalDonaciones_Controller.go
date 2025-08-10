package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTotalDonacionesController struct {
	getTotalDonaciones *application.GetTotalDonaciones
}

func NewGetTotalDonacionesController(getTotalDonaciones *application.GetTotalDonaciones) *GetTotalDonacionesController {
	return &GetTotalDonacionesController{
		getTotalDonaciones: getTotalDonaciones,
	}
}

func (gtdc *GetTotalDonacionesController) Execute(c *gin.Context) {
	total, err := gtdc.getTotalDonaciones.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}
