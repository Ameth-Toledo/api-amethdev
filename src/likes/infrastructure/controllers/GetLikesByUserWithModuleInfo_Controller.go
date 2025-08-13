package controllers

import (
	"AmethToledo/src/likes/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetLikesByUserWithModuleInfoController struct {
	useCase *application.GetLikesByUserWithModuleInfoUseCase
}

func NewGetLikesByUserWithModuleInfoController(useCase *application.GetLikesByUserWithModuleInfoUseCase) *GetLikesByUserWithModuleInfoController {
	return &GetLikesByUserWithModuleInfoController{
		useCase: useCase,
	}
}

func (c *GetLikesByUserWithModuleInfoController) Execute(ctx *gin.Context) {
	usuarioIDStr := ctx.Query("usuario_id")
	if usuarioIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "usuario_id es requerido",
		})
		return
	}

	usuarioID, err := strconv.Atoi(usuarioIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "usuario_id inválido",
		})
		return
	}

	likes, err := c.useCase.Execute(usuarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Error al obtener los likes del usuario con información de módulo",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Likes del usuario obtenidos exitosamente",
		"data": gin.H{
			"usuario_id": usuarioID,
			"likes":      likes,
			"total":      len(likes),
		},
	})
}