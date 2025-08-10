package controllers

import (
	"AmethToledo/src/contenidos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllContenidosController struct {
	getAllContenidos *application.GetAllContenidos
	getContenidoById *application.GetContenidoById
}

func NewGetAllContenidosController(getAllContenidos *application.GetAllContenidos, getContenidoById *application.GetContenidoById) *GetAllContenidosController {
	return &GetAllContenidosController{
		getAllContenidos: getAllContenidos,
		getContenidoById: getContenidoById,
	}
}

func (gac *GetAllContenidosController) Execute(c *gin.Context) {
	// Verificar si hay parámetros de consulta
	id := c.Query("id")

	// Si hay ID, buscar por ID
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		contenido, err := gac.getContenidoById.Execute(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if contenido == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
			return
		}

		c.JSON(http.StatusOK, *contenido)
		return
	}

	// Si no hay parámetros, devolver todos los contenidos
	contenidos, err := gac.getAllContenidos.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contenidos)
}
