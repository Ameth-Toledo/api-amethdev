package controllers

import (
	"AmethToledo/src/comentarios/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllComentariosWithUserController struct {
	getAllComentariosWithUser *application.GetAllComentariosWithUser
}

func NewGetAllComentariosWithUserController(getAllComentariosWithUser *application.GetAllComentariosWithUser) *GetAllComentariosWithUserController {
	return &GetAllComentariosWithUserController{getAllComentariosWithUser: getAllComentariosWithUser}
}

func (gacwuc *GetAllComentariosWithUserController) Execute(c *gin.Context) {
	comentarios, err := gacwuc.getAllComentariosWithUser.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comentarios": comentarios})
}
