package controllers

import (
	"AmethToledo/src/modulos/application"
	"AmethToledo/src/modulos/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateModuloController struct {
	updateModulo *application.UpdateModulo
}

func NewUpdateModuloController(updateModulo *application.UpdateModulo) *UpdateModuloController {
	return &UpdateModuloController{
		updateModulo: updateModulo,
	}
}

type UpdateModuloRequest struct {
	IdCurso       int    `json:"id_curso" binding:"required"`
	ImagenPortada string `json:"imagen_portada" binding:"required"`
	Titulo        string `json:"titulo" binding:"required"`
	Descripcion   string `json:"descripcion" binding:"required"`
}

func (um *UpdateModuloController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateModuloRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modulo := entities.Modulo{
		ID:            id,
		IdCurso:       req.IdCurso,
		ImagenPortada: req.ImagenPortada,
		Titulo:        req.Titulo,
		Descripcion:   req.Descripcion,
	}

	err = um.updateModulo.Execute(modulo)
	if err != nil {
		if err.Error() == "modulo not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Módulo no encontrado"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Módulo actualizado exitosamente",
		"modulo":  modulo,
	})
}
