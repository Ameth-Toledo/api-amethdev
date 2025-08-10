package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeleteComentarioController struct {
	deleteComentario *application.DeleteComentario
}

func NewDeleteComentarioController(deleteComentario *application.DeleteComentario) *DeleteComentarioController {
	return &DeleteComentarioController{deleteComentario: deleteComentario}
}

func (dcc *DeleteComentarioController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dcc.deleteComentario.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentario eliminado exitosamente"})
}
