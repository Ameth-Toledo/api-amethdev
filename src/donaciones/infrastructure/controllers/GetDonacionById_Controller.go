package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetDonacionByIdController struct {
	getDonacionById *application.GetDonacionById
}

func NewGetDonacionByIdController(getDonacionById *application.GetDonacionById) *GetDonacionByIdController {
	return &GetDonacionByIdController{
		getDonacionById: getDonacionById,
	}
}

func (gdc *GetDonacionByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	donacion, err := gdc.getDonacionById.Execute(id)
	if err != nil {
		if err.Error() == "donación no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"donacion": donacion,
	})
}
