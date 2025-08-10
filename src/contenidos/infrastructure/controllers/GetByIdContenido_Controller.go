package controllers

import (
	"AmethToledo/src/contenidos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetContenidoByIdController struct {
	getContenidoById *application.GetContenidoById
}

func NewGetContenidoByIdController(getContenidoById *application.GetContenidoById) *GetContenidoByIdController {
	return &GetContenidoByIdController{
		getContenidoById: getContenidoById,
	}
}

func (gcb *GetContenidoByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	contenido, err := gcb.getContenidoById.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if contenido == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"contenido": contenido,
	})
}
