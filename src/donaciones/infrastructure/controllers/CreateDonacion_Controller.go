package controllers

import (
	"AmethToledo/src/donaciones/application"
	"AmethToledo/src/donaciones/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateDonacionController struct {
	createDonacion *application.CreateDonacion
}

func NewCreateDonacionController(createDonacion *application.CreateDonacion) *CreateDonacionController {
	return &CreateDonacionController{
		createDonacion: createDonacion,
	}
}

type CreateDonacionRequest struct {
	UsuarioID     int     `json:"usuario_id" binding:"required"`
	ModuloID      int     `json:"modulo_id" binding:"required"`
	Monto         float64 `json:"monto" binding:"required,gt=0"`
	Moneda        string  `json:"moneda"`
	Estado        string  `json:"estado" binding:"required"`
	MetodoPago    string  `json:"metodo_pago"`
	TransactionID string  `json:"transaction_id"`
	PaymentID     string  `json:"payment_id"`
}

func (dc *CreateDonacionController) Execute(c *gin.Context) {
	var req CreateDonacionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Establecer moneda por defecto si no se especifica
	if req.Moneda == "" {
		req.Moneda = "MXN"
	}

	donacion := entities.Donacion{
		UsuarioID:     req.UsuarioID,
		ModuloID:      req.ModuloID,
		Monto:         req.Monto,
		Moneda:        req.Moneda,
		Estado:        req.Estado,
		MetodoPago:    req.MetodoPago,
		TransactionID: req.TransactionID,
		PaymentID:     req.PaymentID,
		FechaPago:     time.Now(),
	}

	savedDonacion, err := dc.createDonacion.Execute(donacion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Donación creada exitosamente",
		"donacion": savedDonacion,
	})
}
