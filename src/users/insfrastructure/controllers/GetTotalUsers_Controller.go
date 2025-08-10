package controllers

import (
	"AmethToledo/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetTotalUsersController struct {
	getTotalUsers *application.GetTotalUsers
}

func NewGetTotalUsersController(getTotalUsersController *application.GetTotalUsers) *GetTotalUsersController {
	return &GetTotalUsersController{
		getTotalUsers: getTotalUsersController,
	}
}
func (gtu *GetTotalUsersController) Execute(c *gin.Context) {
	total, err := gtu.getTotalUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"message": "Total de usuarios obtenidos exitosamente",
	})
}
