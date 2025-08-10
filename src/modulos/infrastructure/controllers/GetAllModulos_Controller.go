package controllers

import (
	"AmethToledo/src/modulos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllModulosController struct {
	getAllModulos *application.GetAllModulos
	getModuloById *application.GetModuloById
	searchModulos *application.SearchModulos
}

func NewGetAllModulosController(getAllModulos *application.GetAllModulos, getModuloById *application.GetModuloById, searchModulos *application.SearchModulos) *GetAllModulosController {
	return &GetAllModulosController{
		getAllModulos: getAllModulos,
		getModuloById: getModuloById,
		searchModulos: searchModulos,
	}
}

func (gam *GetAllModulosController) Execute(c *gin.Context) {
	// Verificar si hay parámetros de consulta
	id := c.Query("id")
	titulo := c.Query("titulo")

	// Si hay ID, buscar por ID
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		modulo, err := gam.getModuloById.Execute(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if modulo == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Módulo no encontrado"})
			return
		}

		c.JSON(http.StatusOK, *modulo)
		return
	}

	// Si hay título, buscar por título
	if titulo != "" {
		modulos, err := gam.searchModulos.Execute(titulo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, modulos)
		return
	}

	// Si no hay parámetros, devolver todos los módulos
	modulos, err := gam.getAllModulos.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, modulos)
}
