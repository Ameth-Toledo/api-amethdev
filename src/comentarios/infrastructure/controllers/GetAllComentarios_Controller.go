package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllComentariosController struct {
	getAllComentarios *application.GetAllComentarios
}

func NewGetAllComentariosController(getAllComentarios *application.GetAllComentarios) *GetAllComentariosController {
	return &GetAllComentariosController{getAllComentarios: getAllComentarios}
}

func (gacc *GetAllComentariosController) Execute(c *gin.Context) {
	comentarios, err := gacc.getAllComentarios.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comentarios": comentarios})
}
