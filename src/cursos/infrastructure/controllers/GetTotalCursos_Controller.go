package controllers

import (
	"AmethToledo/src/cursos/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTotalCursosController struct {
	getTotalCursos *application.GetTotalCursos
}

func NewGetTotalCursosController(getTotalCursos *application.GetTotalCursos) *GetTotalCursosController {
	return &GetTotalCursosController{
		getTotalCursos: getTotalCursos,
	}
}

func (gtc *GetTotalCursosController) Execute(c *gin.Context) {
	total, err := gtc.getTotalCursos.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_cursos": total,
		"message":      "Total de cursos obtenido exitosamente",
	})
}
