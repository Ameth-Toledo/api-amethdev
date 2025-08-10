package controllers

import (
	"AmethToledo/src/contenidos/application"
	"AmethToledo/src/contenidos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateContenidoController struct {
	updateContenido *application.UpdateContenido
}

func NewUpdateContenidoController(updateContenido *application.UpdateContenido) *UpdateContenidoController {
	return &UpdateContenidoController{
		updateContenido: updateContenido,
	}
}

type UpdateContenidoRequest struct {
	IdModulo          int    `json:"id_modulo" binding:"required"`
	ImagenPortada     string `json:"imagen_portada" binding:"required"`
	Titulo            string `json:"titulo" binding:"required"`
	Descripcion       string `json:"descripcion" binding:"required"`
	VideoURL          string `json:"video_url"`
	DescripcionModule string `json:"descripcion_module" binding:"required"`
	Repositorio       string `json:"repositorio"`
}

func (uc *UpdateContenidoController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateContenidoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contenido := entities.Contenido{
		ID:                id,
		IdModulo:          req.IdModulo,
		ImagenPortada:     req.ImagenPortada,
		Titulo:            req.Titulo,
		Descripcion:       req.Descripcion,
		VideoURL:          req.VideoURL,
		DescripcionModule: req.DescripcionModule,
		Repositorio:       req.Repositorio, // Puede estar vacío
	}

	err = uc.updateContenido.Execute(contenido)
	if err != nil {
		if err.Error() == "contenido not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Contenido actualizado exitosamente",
		"contenido": contenido,
	})
}
