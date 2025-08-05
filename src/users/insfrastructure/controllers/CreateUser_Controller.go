package controllers

import (
	"AmethToledo/src/users/application"
	"AmethToledo/src/users/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateUserController struct {
	createUser  *application.CreateUser
	authService *application.AuthService
}

func NewCreateUserController(createUser *application.CreateUser, authService *application.AuthService) *CreateUserController {
	return &CreateUserController{
		createUser:  createUser,
		authService: authService,
	}
}

type CreateUserRequest struct {
	Nombres         string `json:"nombres" binding:"required"`
	ApellidoPaterno string `json:"apellido_paterno" binding:"required"`
	ApellidoMaterno string `json:"apellido_materno"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	RolID           int    `json:"rol_id" binding:"required"`
	Avatar          int    `json:"avatar"`
}

func (uc *CreateUserController) Execute(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Avatar < 1 || req.Avatar > 5 {
		req.Avatar = 1
	}

	user := entities.User{
		Nombres:         req.Nombres,
		ApellidoPaterno: req.ApellidoPaterno,
		ApellidoMaterno: req.ApellidoMaterno,
		Email:           req.Email,
		PasswordHash:    req.Password,
		RolID:           req.RolID,
		Avatar:          req.Avatar,
		FechaRegistro:   time.Now(),
	}

	savedUser, err := uc.authService.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario creado exitosamente",
		"user": gin.H{
			"id":               savedUser.ID,
			"nombres":          savedUser.Nombres,
			"apellido_paterno": savedUser.ApellidoPaterno,
			"apellido_materno": savedUser.ApellidoMaterno,
			"email":            savedUser.Email,
			"rol_id":           savedUser.RolID,
			"avatar":           savedUser.Avatar,
			"fecha_registro":   savedUser.FechaRegistro,
		},
	})
}
