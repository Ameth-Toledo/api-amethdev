package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllDonacionesController struct {
	getAllDonaciones *application.GetAllDonaciones
}

func NewGetAllDonacionesController(getAllDonaciones *application.GetAllDonaciones) *GetAllDonacionesController {
	return &GetAllDonacionesController{
		getAllDonaciones: getAllDonaciones,
	}
}

func (gdc *GetAllDonacionesController) Execute(c *gin.Context) {
	donaciones, err := gdc.getAllDonaciones.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"donaciones": donaciones,
		"total":      len(donaciones),
	})
}
