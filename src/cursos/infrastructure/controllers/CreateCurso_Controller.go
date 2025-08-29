package controllers

import (
	"AmethToledo/src/cursos/application"
	"AmethToledo/src/cursos/domain/entities"
	"AmethToledo/src/notifications"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateCursoController struct {
	createCurso *application.CreateCurso
	hub         *notifications.Hub
}

func NewCreateCursoController(createCurso *application.CreateCurso, hub *notifications.Hub) *CreateCursoController {
	return &CreateCursoController{
		createCurso: createCurso,
		hub:         hub,
	}
}

type CreateCursoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Nivel       string `json:"nivel" binding:"required"`
	Duracion    string `json:"duracion" binding:"required"`
	Tecnologia  string `json:"tecnologia" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Imagen      string `json:"imagen" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

func (cc *CreateCursoController) Execute(c *gin.Context) {
	var req CreateCursoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curso := entities.Curso{
		Nombre:      req.Nombre,
		Nivel:       req.Nivel,
		Duracion:    req.Duracion,
		Tecnologia:  req.Tecnologia,
		Fecha:       req.Fecha,
		Imagen:      req.Imagen,
		Descripcion: req.Descripcion,
	}

	savedCurso, err := cc.createCurso.Execute(curso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cc.hub.BroadcastNotification(
		"course_created",
		"¡Nuevo curso disponible!",
		map[string]interface{}{
			"curso":   savedCurso,
			"titulo":  savedCurso.Nombre,
			"mensaje": "Se ha agregado un nuevo curso de " + savedCurso.Tecnologia,
		},
	)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Curso creado exitosamente",
		"curso":   savedCurso,
	})
}
