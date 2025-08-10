package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetComentariosByUsuarioController struct {
	getComentariosByUsuario *application.GetComentariosByUsuario
}

func NewGetComentariosByUsuarioController(getComentariosByUsuario *application.GetComentariosByUsuario) *GetComentariosByUsuarioController {
	return &GetComentariosByUsuarioController{getComentariosByUsuario: getComentariosByUsuario}
}

func (gcbuc *GetComentariosByUsuarioController) Execute(c *gin.Context) {
	usuarioIdStr := c.Param("usuarioId")
	usuarioId, err := strconv.Atoi(usuarioIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del usuario inválido"})
		return
	}

	comentarios, err := gcbuc.getComentariosByUsuario.Execute(usuarioId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comentarios": comentarios})
}
