package controllers

import (
	"AmethToledo/src/comentarios/application"
	"AmethToledo/src/comentarios/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateComentarioController struct {
	updateComentario *application.UpdateComentario
}

func NewUpdateComentarioController(updateComentario *application.UpdateComentario) *UpdateComentarioController {
	return &UpdateComentarioController{updateComentario: updateComentario}
}

type UpdateComentarioRequest struct {
	ModuloID  int    `json:"modulo_id" binding:"required"`
	UsuarioID int    `json:"usuario_id" binding:"required"`
	Texto     string `json:"texto" binding:"required"`
}

func (ucc *UpdateComentarioController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateComentarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comentario := entities.Comentario{
		ID:        id,
		ModuloID:  req.ModuloID,
		UsuarioID: req.UsuarioID,
		Texto:     req.Texto,
	}

	err = ucc.updateComentario.Execute(comentario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comentario actualizado exitosamente"})
}
