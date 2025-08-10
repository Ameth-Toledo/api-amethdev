package controllers

import (
	"AmethToledo/src/cursos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteCursoController struct {
	deleteCurso *application.DeleteCurso
}

func NewDeleteCursoController(deleteCurso *application.DeleteCurso) *DeleteCursoController {
	return &DeleteCursoController{
		deleteCurso: deleteCurso,
	}
}

func (dc *DeleteCursoController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dc.deleteCurso.Execute(id)
	if err != nil {
		if err.Error() == "curso not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Curso no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Curso eliminado exitosamente",
	})
}
