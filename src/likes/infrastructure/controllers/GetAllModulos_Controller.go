package controllers

import (
	"AmethToledo/src/likes/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllModulosController struct {
	useCase *application.GetAllModulosUseCase
}

func NewGetAllModulosController(useCase *application.GetAllModulosUseCase) *GetAllModulosController {
	return &GetAllModulosController{
		useCase: useCase,
	}
}

func (c *GetAllModulosController) Execute(ctx *gin.Context) {
	modulos, err := c.useCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error al obtener los módulos",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Módulos obtenidos exitosamente",
		"data": gin.H{
			"modulos": modulos,
			"total":   len(modulos),
		},
	})
}