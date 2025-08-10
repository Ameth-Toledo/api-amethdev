package controllers

import (
	"AmethToledo/src/donaciones/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetDonacionesByUsuarioController struct {
	getDonacionesByUsuario *application.GetDonacionesByUsuario
}

func NewGetDonacionesByUsuarioController(getDonacionesByUsuario *application.GetDonacionesByUsuario) *GetDonacionesByUsuarioController {
	return &GetDonacionesByUsuarioController{
		getDonacionesByUsuario: getDonacionesByUsuario,
	}
}

func (gdc *GetDonacionesByUsuarioController) Execute(c *gin.Context) {
	usuarioIDParam := c.Param("usuario_id")
	usuarioID, err := strconv.Atoi(usuarioIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "usuario_id inválido"})
		return
	}

	donaciones, err := gdc.getDonacionesByUsuario.Execute(usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"donaciones": donaciones,
		"total":      len(donaciones),
		"usuario_id": usuarioID,
	})
}
