package controllers

import (
	"AmethToledo/src/contenidos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteContenidoController struct {
	deleteContenido *application.DeleteContenido
}

func NewDeleteContenidoController(deleteContenido *application.DeleteContenido) *DeleteContenidoController {
	return &DeleteContenidoController{
		deleteContenido: deleteContenido,
	}
}

func (dc *DeleteContenidoController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dc.deleteContenido.Execute(id)
	if err != nil {
		if err.Error() == "contenido not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Contenido eliminado exitosamente",
	})
}
