package controllers

import (
	"AmethToledo/src/modulos/application"
	"AmethToledo/src/modulos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateModuloController struct {
	createModulo *application.CreateModulo
}

func NewCreateModuloController(createModulo *application.CreateModulo) *CreateModuloController {
	return &CreateModuloController{
		createModulo: createModulo,
	}
}

type CreateModuloRequest struct {
	IdCurso       int    `json:"id_curso" binding:"required"`
	ImagenPortada string `json:"imagen_portada" binding:"required"`
	Titulo        string `json:"titulo" binding:"required"`
	Descripcion   string `json:"descripcion" binding:"required"`
}

func (cm *CreateModuloController) Execute(c *gin.Context) {
	var req CreateModuloRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modulo := entities.Modulo{
		IdCurso:       req.IdCurso,
		ImagenPortada: req.ImagenPortada,
		Titulo:        req.Titulo,
		Descripcion:   req.Descripcion,
	}

	savedModulo, err := cm.createModulo.Execute(modulo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Módulo creado exitosamente",
		"modulo":  savedModulo,
	})
}
