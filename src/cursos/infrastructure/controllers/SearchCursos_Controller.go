package controllers

import (
	"AmethToledo/src/cursos/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchCursosController struct {
	searchCursos *application.SearchCursos
}

func NewSearchCursosController(searchCursos *application.SearchCursos) *SearchCursosController {
	return &SearchCursosController{
		searchCursos: searchCursos,
	}
}

func (sc *SearchCursosController) Execute(c *gin.Context) {
	nombre := c.Query("nombre")

	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'nombre' es requerido"})
		return
	}

	cursos, err := sc.searchCursos.Execute(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cursos":   cursos,
		"total":    len(cursos),
		"busqueda": nombre,
	})
}
