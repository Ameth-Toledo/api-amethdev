package controllers

import (
	"AmethToledo/src/cursos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetCursoByIdController struct {
	getCursoById *application.GetCursoById
}

func NewGetCursoByIdController(getCursoById *application.GetCursoById) *GetCursoByIdController {
	return &GetCursoByIdController{
		getCursoById: getCursoById,
	}
}

func (gcb *GetCursoByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	curso, err := gcb.getCursoById.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if curso == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"curso": curso,
	})
}
