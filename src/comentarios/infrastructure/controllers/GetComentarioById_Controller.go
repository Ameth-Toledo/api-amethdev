package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetComentarioByIdController struct {
	getComentarioById *application.GetComentarioById
}

func NewGetComentarioByIdController(getComentarioById *application.GetComentarioById) *GetComentarioByIdController {
	return &GetComentarioByIdController{getComentarioById: getComentarioById}
}

func (gcbic *GetComentarioByIdController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	comentario, err := gcbic.getComentarioById.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if comentario == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comentario": comentario})
}
