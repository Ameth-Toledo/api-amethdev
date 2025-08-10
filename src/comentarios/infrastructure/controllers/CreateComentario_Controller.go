package controllers

import (
	"AmethToledo/src/comentarios/application"
	"AmethToledo/src/comentarios/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateComentarioController struct {
	createComentario *application.CreateComentario
}

func NewCreateComentarioController(createComentario *application.CreateComentario) *CreateComentarioController {
	return &CreateComentarioController{createComentario: createComentario}
}

type CreateComentarioRequest struct {
	ModuloID  int    `json:"modulo_id" binding:"required"`
	UsuarioID int    `json:"usuario_id" binding:"required"`
	Texto     string `json:"texto" binding:"required"`
}

func (ccc *CreateComentarioController) Execute(c *gin.Context) {
	var req CreateComentarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comentario := entities.Comentario{
		ModuloID:  req.ModuloID,
		UsuarioID: req.UsuarioID,
		Texto:     req.Texto,
		Fecha:     time.Now(),
	}

	savedComentario, err := ccc.createComentario.Execute(comentario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":    "Comentario creado exitosamente",
		"comentario": savedComentario,
	})
}
