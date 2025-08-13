package controllers

import (
	"AmethToledo/src/likes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLikesByModuloWithUserInfoController struct {
	useCase *application.GetLikesByModuloWithUserInfoUseCase
}

func NewGetLikesByModuloWithUserInfoController(useCase *application.GetLikesByModuloWithUserInfoUseCase) *GetLikesByModuloWithUserInfoController {
	return &GetLikesByModuloWithUserInfoController{
		useCase: useCase,
	}
}

func (c *GetLikesByModuloWithUserInfoController) Execute(ctx *gin.Context) {
	moduloIDStr := ctx.Param("modulo_id")
	moduloID, err := strconv.Atoi(moduloIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "ID de módulo inválido",
		})
		return
	}

	likes, err := c.useCase.Execute(moduloID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error al obtener los likes con información de usuario",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Likes obtenidos exitosamente",
		"data": gin.H{
			"modulo_id": moduloID,
			"likes":     likes,
			"total":     len(likes),
		},
	})
}