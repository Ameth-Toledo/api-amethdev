package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteDonacionController struct {
	deleteDonacion *application.DeleteDonacion
}

func NewDeleteDonacionController(deleteDonacion *application.DeleteDonacion) *DeleteDonacionController {
	return &DeleteDonacionController{
		deleteDonacion: deleteDonacion,
	}
}

func (ddc *DeleteDonacionController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ddc.deleteDonacion.Execute(id)
	if err != nil {
		if err.Error() == "donación no encontrada" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Donación eliminada exitosamente",
	})
}
