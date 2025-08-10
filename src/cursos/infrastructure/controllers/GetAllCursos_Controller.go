package controllers

import (
	"AmethToledo/src/cursos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllCursosController struct {
	getAllCursos *application.GetAllCursos
	getCursoById *application.GetCursoById
	searchCursos *application.SearchCursos
}

func NewGetAllCursosController(getAllCursos *application.GetAllCursos, getCursoById *application.GetCursoById, searchCursos *application.SearchCursos) *GetAllCursosController {
	return &GetAllCursosController{
		getAllCursos: getAllCursos,
		getCursoById: getCursoById,
		searchCursos: searchCursos,
	}
}

func (gac *GetAllCursosController) Execute(c *gin.Context) {
	// Verificar si hay parámetros de consulta
	id := c.Query("id")
	nombre := c.Query("nombre")

	// Si hay ID, buscar por ID
	if id != "" {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		curso, err := gac.getCursoById.Execute(idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if curso == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
			return
		}

		c.JSON(http.StatusOK, *curso)
		return
	}

	// Si hay nombre, buscar por nombre
	if nombre != "" {
		cursos, err := gac.searchCursos.Execute(nombre)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, cursos)
		return
	}

	// Si no hay parámetros, devolver todos los cursos
	cursos, err := gac.getAllCursos.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cursos)
}
