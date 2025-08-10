package controllers

import (
	"AmethToledo/src/cursos/application"
	"AmethToledo/src/cursos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateCursoController struct {
	updateCurso *application.UpdateCurso
}

func NewUpdateCursoController(updateCurso *application.UpdateCurso) *UpdateCursoController {
	return &UpdateCursoController{
		updateCurso: updateCurso,
	}
}

type UpdateCursoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Nivel       string `json:"nivel" binding:"required"`
	Duracion    string `json:"duracion" binding:"required"`
	Tecnologia  string `json:"tecnologia" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Imagen      string `json:"imagen" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

func (uc *UpdateCursoController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateCursoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curso := entities.Curso{
		ID:          id,
		Nombre:      req.Nombre,
		Nivel:       req.Nivel,
		Duracion:    req.Duracion,
		Tecnologia:  req.Tecnologia,
		Fecha:       req.Fecha,
		Imagen:      req.Imagen,
		Descripcion: req.Descripcion,
	}

	err = uc.updateCurso.Execute(curso)
	if err != nil {
		if err.Error() == "curso not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Curso actualizado exitosamente",
		"curso":   curso,
	})
}
