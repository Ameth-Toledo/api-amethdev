package controllers

import (
	"AmethToledo/src/modulos/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetModuloByIdController struct {
	getModuloById *application.GetModuloById
}

func NewGetModuloByIdController(getModuloById *application.GetModuloById) *GetModuloByIdController {
	return &GetModuloByIdController{
		getModuloById: getModuloById,
	}
}

func (gmb *GetModuloByIdController) Execute(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	modulo, err := gmb.getModuloById.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if modulo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Módulo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"modulo": modulo,
	})
}
