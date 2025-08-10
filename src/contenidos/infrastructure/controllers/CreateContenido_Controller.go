package controllers

import (
	"AmethToledo/src/contenidos/application"
	"AmethToledo/src/contenidos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateContenidoController struct {
	createContenido *application.CreateContenido
}

func NewCreateContenidoController(createContenido *application.CreateContenido) *CreateContenidoController {
	return &CreateContenidoController{
		createContenido: createContenido,
	}
}

type CreateContenidoRequest struct {
	IdModulo          int    `json:"id_modulo" binding:"required"`
	ImagenPortada     string `json:"imagen_portada" binding:"required"`
	Titulo            string `json:"titulo" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
	VideoURL          string `json:"video_url"`
	DescripcionModule string `json:"descripcion_module" binding:"required"`
	Repositorio       string `json:"repositorio"`
}

func (cc *CreateContenidoController) Execute(c *gin.Context) {
	var req CreateContenidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contenido := entities.Contenido{
		IdModulo:          req.IdModulo,
		ImagenPortada:     req.ImagenPortada,
		Titulo:            req.Titulo,
		Descripcion:       req.Descripcion,
		VideoURL:          req.VideoURL,
		DescripcionModule: req.DescripcionModule,
		Repositorio:       req.Repositorio, // Puede estar vacío
	}

	savedContenido, err := cc.createContenido.Execute(contenido)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Contenido creado exitosamente",
		"contenido": savedContenido,
	})
}
