package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetDonacionesByModuloController struct {
	getDonacionesByModulo *application.GetDonacionesByModulo
}

func NewGetDonacionesByModuloController(getDonacionesByModulo *application.GetDonacionesByModulo) *GetDonacionesByModuloController {
	return &GetDonacionesByModuloController{
		getDonacionesByModulo: getDonacionesByModulo,
	}
}

func (gdc *GetDonacionesByModuloController) Execute(c *gin.Context) {
	moduloIDParam := c.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "modulo_id inválido"})
		return
	}

	donaciones, err := gdc.getDonacionesByModulo.Execute(moduloID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"donaciones": donaciones,
		"total":      len(donaciones),
		"modulo_id":  moduloID,
	})
}
