package controllers

import (
	"AmethToledo/src/users/application"
	"AmethToledo/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateUserController struct {
	updateUser *application.UpdateUser
}

func NewUpdateUserController(updateUser *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{updateUser: updateUser}
}

type UpdateUserRequest struct {
	Nombres         string `json:"nombres" binding:"required"`
	ApellidoPaterno string `json:"apellido_paterno" binding:"required"`
	ApellidoMaterno string `json:"apellido_materno"`
	Email           string `json:"email" binding:"required,email"`
	RolID           int    `json:"rol_id" binding:"required"`
	Avatar          int    `json:"avatar"`
}

func (uc *UpdateUserController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar avatar (debe estar entre 1 y 5)
	if req.Avatar < 1 || req.Avatar > 5 {
		req.Avatar = 1 // Valor por defecto
	}

	user := entities.User{
		ID:              id,
		Nombres:         req.Nombres,
		ApellidoPaterno: req.ApellidoPaterno,
		ApellidoMaterno: req.ApellidoMaterno,
		Email:           req.Email,
		RolID:           req.RolID,
		Avatar:          req.Avatar,
	}

	err = uc.updateUser.Execute(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}
