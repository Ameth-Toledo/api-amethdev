package controllers

import (
	"AmethToledo/src/modulos/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchModulosController struct {
	searchModulos *application.SearchModulos
}

func NewSearchModulosController(searchModulos *application.SearchModulos) *SearchModulosController {
	return &SearchModulosController{
		searchModulos: searchModulos,
	}
}

func (sm *SearchModulosController) Execute(c *gin.Context) {
	titulo := c.Query("titulo")

	if titulo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'titulo' es requerido"})
		return
	}

	modulos, err := sm.searchModulos.Execute(titulo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"modulos":  modulos,
		"total":    len(modulos),
		"busqueda": titulo,
	})
}
