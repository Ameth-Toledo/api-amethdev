package controllers

import (
	"AmethToledo/src/modulos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteModuloController struct {
	deleteModulo *application.DeleteModulo
}

func NewDeleteModuloController(deleteModulo *application.DeleteModulo) *DeleteModuloController {
	return &DeleteModuloController{
		deleteModulo: deleteModulo,
	}
}

func (dm *DeleteModuloController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dm.deleteModulo.Execute(id)
	if err != nil {
		if err.Error() == "modulo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Módulo no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Módulo eliminado exitosamente",
	})
}
