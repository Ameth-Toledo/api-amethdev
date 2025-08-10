package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTotalComentariosController struct {
	getTotalComentarios *application.GetTotalComentarios
}

func NewGetTotalComentariosController(getTotalComentarios *application.GetTotalComentarios) *GetTotalComentariosController {
	return &GetTotalComentariosController{getTotalComentarios: getTotalComentarios}
}

func (gtcc *GetTotalComentariosController) Execute(c *gin.Context) {
	total, err := gtcc.getTotalComentarios.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total": total})
}
